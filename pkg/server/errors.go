package server

import "net/http"

type ErrorBody struct {
	Message string `json:"message"`
}

func NewBadRequestError(err error) *Response {
	return NewError(http.StatusBadRequest, err)
}

func NewError(statusCode int, err error) *Response {
	return &Response{
		StatusCode: statusCode,
		Body: &ErrorBody{
			Message: err.Error(),
		},
	}
}

func NewForbiddenError(err error) *Response {
	return NewError(http.StatusForbidden, err)
}

func NewInternalServerError(err error) *Response {
	return NewError(http.StatusInternalServerError, err)
}

func NewNotFoundError(err error) *Response {
	return NewError(http.StatusNotFound, err)
}

func NewServiceUnavailableError(err error) *Response {
	return NewError(http.StatusServiceUnavailable, err)
}

func NewUnauthorizedError(err error) *Response {
	return NewError(http.StatusUnauthorized, err)
}
