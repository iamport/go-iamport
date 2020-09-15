package util

import (
	"errors"
	"net/http"
)

const (
	CodeOK = 0

	ErrStatusUnauthorized = "iamport: unauthorized"
	ErrStatusNotFound     = "iamport: invalid imp_uid"
	ErrUnknown            = "iamport: unknown error"
)

func ErrorHandler(res *http.Response) error {
	switch res.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusUnauthorized:
		return errors.New(ErrStatusUnauthorized)
	case http.StatusNotFound:
		return errors.New(ErrStatusNotFound)
	default:
		return errors.New(ErrUnknown)
	}
}
