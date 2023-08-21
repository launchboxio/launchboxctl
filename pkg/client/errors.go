package client

import "fmt"

type ApiError struct {
	Status  int
	Message string
}

func (a ApiError) Error() string {
	return fmt.Sprintf("%d: %s", a.Status, a.Message)
}

func NewUnauthorizedError() ApiError {
	return ApiError{401, "Unauthorized"}
}

func NewForbiddenError() ApiError {
	return ApiError{403, "Forbidden"}
}

func NewNotFoundError() ApiError {
	return ApiError{404, "Not Found"}
}

func NewInternalServerError() ApiError {
	return ApiError{500, "An unexpected error occured"}
}
