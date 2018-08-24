package errors

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// Type .
type Type string

// Types
const (
	Data          Type = "data"
	Action        Type = "action"
	NotFound      Type = "not-found"
	Remote        Type = "remote"
	Authorization Type = "authorization"
)

// Error .
type Error struct {
	Type        Type
	SourceError error
}

// Error .
func (e Error) Error() string { return e.SourceError.Error() }

// NewData .
func NewData(message interface{}) error {
	return Error{
		Type:        Data,
		SourceError: errors.WithStack(fmt.Errorf("%s", message)),
	}
}

// NewAction .
func NewAction(message interface{}) error {
	return Error{
		Type:        Action,
		SourceError: errors.WithStack(fmt.Errorf("%s", message)),
	}
}

// NewNotFound .
func NewNotFound() error {
	return Error{
		Type:        NotFound,
		SourceError: errors.WithStack(errors.New(string(NotFound))),
	}
}

// NewRemote .
func NewRemote() error {
	return Error{
		Type: Remote,
		SourceError: errors.WithStack(errors.
			New(http.StatusText(http.StatusServiceUnavailable))),
	}
}

// NewAuthorization .
func NewAuthorization(message interface{}) error {
	return Error{
		Type:        Authorization,
		SourceError: errors.WithStack(fmt.Errorf("%s", message)),
	}
}
