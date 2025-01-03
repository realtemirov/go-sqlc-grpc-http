package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/realtemirov/go-sqlc-grpc-http/config"
	db "github.com/realtemirov/go-sqlc-grpc-http/db/sqlc"
)

type Handler struct {
	cfg  *config.Config
	repo db.Store
}

type Response struct {
	Data interface{} `json:"data"`
}

func NewHandler(cfg *config.Config, repo db.Store) *Handler {
	handlers := Handler{
		cfg:  cfg,
		repo: repo,
	}

	return &handlers
}

//nolint:unused // This is a helper function for the handlers
func (h *Handler) handleResponse(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, Response{
		Data: data,
	})
}

var (
	ErrBadRequest       = errors.New("bad request")
	ErrInternal         = errors.New("internal error")
	ErrWrongCredentials = errors.New("wrong credentials")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrForbidden        = errors.New("forbidden")
)

//nolint:unused // This is a helper function for the handlers
func (h *Handler) handleError(c *gin.Context, err error) {
	var statusCode int

	switch {
	case errors.Is(err, ErrBadRequest):
		statusCode = http.StatusBadRequest
	case errors.Is(err, ErrInternal):
		statusCode = http.StatusInternalServerError
	case errors.Is(err, ErrWrongCredentials):
		statusCode = http.StatusUnauthorized
	case errors.Is(err, ErrUnauthorized):
		statusCode = http.StatusUnauthorized
	case errors.Is(err, ErrForbidden):
		statusCode = http.StatusForbidden
	default:
		statusCode = http.StatusInternalServerError
	}

	c.JSON(statusCode, Response{
		Data: err.Error(),
	})
}

func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "pong",
	})
}
