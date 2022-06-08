package model

import (
	"net/http"

	"github.com/pkg/errors"
)

const (
	CodeErrUnexpected         = 0
	CodeErrDiscountNotFound   = 1
	CodeErrDiscountUsed       = 2
	CodeErrServiceUnavailable = 3
)

var (
	ErrUnexpected         = errors.New("unexpected error")
	ErrDiscountNotFound   = errors.New("discount code not found")
	ErrDiscountUsed       = errors.New("discount used")
	ErrServiceUnavailable = errors.New("service unavailable")
)

func ErrToCode(err error) int {
	switch errors.Cause(err) {
	case ErrDiscountNotFound:
		return CodeErrDiscountNotFound
	case ErrDiscountUsed:
		return CodeErrDiscountUsed
	default:
		return CodeErrUnexpected
	}
}

func ErrToHTTPStatus(err error) int {
	switch errors.Cause(err) {
	case ErrDiscountNotFound, ErrDiscountUsed:
		return http.StatusNotFound
	case ErrServiceUnavailable:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
