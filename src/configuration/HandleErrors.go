package handleError

import "net/http"

type HandleError struct {
	Message string  `json:"message"`
	Err     string  `json:"error"`
	Code    int     `json:"code"`
	Cases   []Cases `json:"cases"`
}

type Cases struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewHandleErr(message, err string, code int, cases []Cases) *HandleError {
	return &HandleError{
		Message: message,
		Err:     err,
		Code:    code,
		Cases:   cases,
	}
}

func (r *HandleError) Error() string {
	return r.Message
}

func NewBadRequestError(message string) *HandleError {
	return &HandleError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewUnauthorizedRequestError(message string) *HandleError {
	return &HandleError{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func NewBadRequestValidationError(message string, causes []Cases) *HandleError {
	return &HandleError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Cases:   causes,
	}
}

func NewInternalServerError(message string) *HandleError {
	return &HandleError{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *HandleError {
	return &HandleError{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *HandleError {
	return &HandleError{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}
