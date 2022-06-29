package httperrors

import (
	"log"
	"net/http"
)

// HTTPError define data struct for HTTP Errors
type HTTPError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Error      string `json:"error"`
}

// NewHTTPError creates and return instance of HTTPError
func NewHTTPError(statusCode int, msg, err string) *HTTPError {
	return &HTTPError{
		StatusCode: statusCode,
		Message:    msg,
		Error:      err,
	}
}

// NewInternalServerError creates an error with InternalServerError info
func NewInternalServerError(msg string, err error) *HTTPError {
	log.Println(err.Error())
	return NewHTTPError(http.StatusInternalServerError, msg, "InternalServerError")
}

// NewUnexpectedError creates an error with InternalServerError info and defined message
func NewUnexpectedError(err error) *HTTPError {
	return NewInternalServerError("something went wrong!", err)
}

// NewBadRequestError creates an error with BadRequest info
func NewBadRequestError(msg string) *HTTPError {
	return NewHTTPError(http.StatusBadRequest, msg, "BadRequest")
}

// NewNotFoundError creates an error with NotFound info
func NewNotFoundError(msg string) *HTTPError {
	return NewHTTPError(http.StatusNotFound, msg, "NotFound")
}

// NewUnauthorizedError creates an error with Unauthorized info
func NewUnauthorizedError(msg string) *HTTPError {
	return NewHTTPError(http.StatusUnauthorized, msg, "Unauthorized")
}
