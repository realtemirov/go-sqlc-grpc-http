package xerrors

import "google.golang.org/grpc/codes"

const (
	ValidationFailed            = "validation_failed"
	DBOperationFailed           = "db_failed"
	CachingFailed               = "caching_failed"
	QueueFailed                 = "queue_failed"
	InternalError               = "internal_error"
	NotFound                    = "not_found"
	Unauthorized                = "unauthorized"
	Forbidden                   = "forbidden"
	AlreadyExists               = "already_exists"
	InvalidToken                = "invalid_token"
	RateLimitReached            = "rate_limit_reached"
	UserBlocked                 = "user_blocked"
	ThirdPartyIntegrationFailed = "third_party_integration_failed"
	DateConversionFailed        = "date_conversion_failed"
	NoDataForPeriod             = "no_data_for_period"

	InternalErrorCode = 500001
)

type Message struct {
	Message   string     `json:"message" csv:"message"`
	Code      codes.Code `json:"code" csv:"code"`
	Labels    Labels     `json:"labels" csv:"labels"`
	ErrorCode int32      `json:"error_code" csv:"error_code"`
}

type Labels struct {
	Uz string `json:"uz" csv:"uz"`
	Ru string `json:"ru" csv:"ru"`
	En string `json:"en" csv:"en"`
}

var (
	//nolint:gochecknoglobals //errors are global
	GlobalErrors = map[string]map[string]Message{
		InternalError: {
			NotFound: {
				Message:   "Internal error",
				Code:      codes.Internal,
				ErrorCode: InternalErrorCode,
				Labels: Labels{
					Uz: "Serverda ichki xatolik yuzaga keldi",
					Ru: "Произошла внутренняя ошибка сервера",
					En: "Internal server error happened",
				},
			},
		},
	}
)
