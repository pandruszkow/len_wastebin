package netshare

import (
	"errors"
)

var (
	ErrBadRequest       = errors.New("bad request")           // 400
	ErrNotFound         = errors.New("not found")             // 404
	ErrMethodNotAllowed = errors.New("method not allowed")    // 405
	ErrInternal         = errors.New("internal server error") // 500
)
