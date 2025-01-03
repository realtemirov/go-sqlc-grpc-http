package jwt

import (
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/realtemirov/go-sqlc-grpc-http/pkg/libraries/passcode"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
)

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
	PassCode     TokenType = "passcode"
)

type Cred struct {
	Secret string
	Expire time.Duration
}

type Jwt struct {
	Creds map[TokenType]Cred
}

type Authenticator interface {
	GenerateJWT(mapToken map[string]interface{}, tokenType TokenType) (*TokenResponse, error)
	CheckPasscode(params CheckClaimsParams) error
	GetAuthInfo(c *gin.Context) (*AuthResp, error)
	ParseToken(tokenString string, tokenType TokenType) (*AuthResp, error)
	GetExpireTime(tokenType TokenType) time.Time
}

type CheckClaimsParams struct {
	Token string
	Code  string
	Phone string
	Email string
	Dev   bool
}

type AuthResp struct {
	ID          string
	Role        int32
	SessionID   string
	SignedToken string
}

type TokenResponse struct {
	Token     string
	ExpiresAt int64
}

func New(params map[TokenType]Cred) Authenticator {
	return &Jwt{
		Creds: params,
	}
}

// GenerateToken generates a new JWT token string - tokenType - [access, passcode, refresh].
func (j *Jwt) GenerateJWT(mapToken map[string]interface{}, tokenType TokenType) (*TokenResponse, error) {
	var (
		token       = jwt.New(jwt.SigningMethodHS256)
		tokenString string
		err         error
	)

	creds, ok := j.Creds[tokenType]
	if !ok {
		return nil, errors.New("invalid token type")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	for key, value := range mapToken {
		claims[key] = value
	}

	claims["iss"] = creds.Secret
	claims["aud"] = creds.Secret
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(creds.Expire).Unix()

	tokenString, err = token.SignedString([]byte(creds.Secret))
	if err != nil {
		return nil, err
	}

	return &TokenResponse{
		Token:     tokenString,
		ExpiresAt: time.Now().Add(creds.Expire).Unix(),
	}, nil
}

func (j *Jwt) extractClaims(tokenString string, tokenType TokenType) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	creds, ok := j.Creds[tokenType]
	if !ok {
		return nil, errors.New("invalid token type")
	}

	token, err = jwt.Parse(tokenString, func(_ *jwt.Token) (interface{}, error) {
		return []byte(creds.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func (j *Jwt) CheckPasscode(params CheckClaimsParams) error {
	claims, err := j.extractClaims(params.Token, PassCode)
	if err != nil {
		return err
	}

	authType := ""
	login := ""
	if params.Email != "" {
		authType = "email"
		login = params.Email
	} else if params.Phone != "" {
		authType = "phone"
		login = params.Phone
	}

	passCodeToken, ok := claims["hashed_code"].(string)
	if !ok {
		return errors.New("invalid token")
	}

	input := passcode.Passcode(login, params.Code, authType)

	err = bcrypt.CompareHashAndPassword([]byte(passCodeToken), []byte(input))
	if err != nil {
		return err
	}

	return nil
}

func (j *Jwt) GetAuthInfo(ctx *gin.Context) (*AuthResp, error) {
	var (
		auth AuthResp
	)

	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		return &auth, nil
	}

	claims, err := j.extractClaims(tokenString, AccessToken)
	if err != nil {
		return nil, err
	}

	id, ok := claims["id"].(string)
	if !ok {
		return nil, errors.New("invalid id")
	}

	sessionID, ok := claims["session_id"].(string)
	if !ok {
		log.Println("invalid session id")
	}

	auth.ID = id
	auth.Role = cast.ToInt32(claims["role"])
	auth.SessionID = sessionID

	return &auth, nil
}

func (j *Jwt) ParseToken(tokenString string, tokenType TokenType) (*AuthResp, error) {
	var (
		auth AuthResp
	)

	claims, err := j.extractClaims(tokenString, tokenType)
	if err != nil {
		return nil, err
	}

	id, ok := claims["id"].(string)
	if !ok {
		return nil, errors.New("invalid id")
	}

	sessionID, ok := claims["session_id"].(string)
	if !ok {
		return nil, errors.New("invalid session id")
	}

	auth.ID = id
	auth.Role = cast.ToInt32(claims["role"])
	auth.SessionID = sessionID

	return &auth, nil
}

func (j *Jwt) GetExpireTime(tokenType TokenType) time.Time {
	creds, ok := j.Creds[tokenType]
	if !ok {
		return time.Now()
	}

	return time.Now().Add(creds.Expire)
}
