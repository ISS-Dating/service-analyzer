package service

import (
	"errors"
)

var (
	ErrorForbidden = errors.New("Forbidden")
)

type Interface interface {
	UpdateWithMessage(message string)
}
