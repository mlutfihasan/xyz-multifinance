package exception

import (
	"errors"
)

var (
	ErrPermissionDenied    = errors.New("permission denied")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrRefreshTokenExpired = errors.New("refresh token expired")
)

type ErrorSendToResponse struct {
	Err string
}

func (e *ErrorSendToResponse) Error() string {
	return e.Err
}
