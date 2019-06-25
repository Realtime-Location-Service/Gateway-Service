package model

import (
	"io"
	"net/http"
)

// HTTPResponse ...
type HTTPResponse struct {
	StatusCode int
	// must be closed after read
	Body io.ReadCloser

	Header http.Header

	Err error
}
