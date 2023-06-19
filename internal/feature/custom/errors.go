package customerror

import "net/http"

type (
	EntityNotFoundError string
	ValidationError     string
	LockedError         string
	UnprocessableEntity string
	Error               string
	ForbiddenError      string
	InternalServerError string

	APIError struct {
		ErrorStr string    `json:"error"`
		Status   int       `json:"status"`
		Cause    CauseList `json:"cause"`
	}

	CauseList []interface{}

	InternalUsageErrors interface {
		EntityNotFoundError | ValidationError | LockedError | UnprocessableEntity | Error
	}
)

func (e EntityNotFoundError) Error() string { return string(e) }
func (e InternalServerError) Error() string { return string(e) }
func (e ForbiddenError) Error() string      { return string(e) }
func (e ValidationError) Error() string     { return string(e) }
func (e LockedError) Error() string         { return string(e) }
func (e APIError) Error() string            { return e.ErrorStr }
func (e UnprocessableEntity) Error() string { return string(e) }
func (e Error) Error() string               { return string(e) }

func NotFoundAPIError(message string) APIError {
	return APIError{"not_found", http.StatusNotFound, CauseList{message}}
}

func ForbiddenAPIError(message string) APIError {
	return APIError{"forbidden", http.StatusForbidden, CauseList{message}}
}

func BadRequestAPIError(message string) APIError {
	return APIError{"bad_request", http.StatusBadRequest, CauseList{message}}
}

func LockedAPIError(message string) APIError {
	return APIError{"locked", http.StatusLocked, CauseList{message}}
}

func TooManyRequestsAPIError(message string) APIError {
	return APIError{"too_many_requests", http.StatusTooManyRequests, CauseList{message}}
}

func UnprocessableEntityAPIError(message string) APIError {
	return APIError{"unprocessable_entity", http.StatusUnprocessableEntity, CauseList{message}}
}

func InternalServerAPIError(err string) APIError {
	return APIError{
		ErrorStr: "internal_server_error",
		Status:   http.StatusInternalServerError,
		Cause:    CauseList{err},
	}
}
