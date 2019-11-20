package response

const (
	StatusOK                   = "ok"
	StatusInternalError        = "internal_error"
	StatusInvalidRequestParams = "invalid_request_params"
	StatusForbidden            = "forbidden"
	StatusUnauthorized         = "unauthorized"
	StatusNotFound             = "not_found"
	StatusNotModified          = "not_modified"
	StatusRevisionTooOld       = "revision_too_old"
	StatusTooManyRequests      = "too_many_requests"
	StatusTooManyAttempts      = "too_many_attempts"
	StatusConflict             = "conflict"
)

type Basic struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}
