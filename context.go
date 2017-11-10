package tuu

import (
	"context"
	"net/http"
)

type Context interface {
	context.Context
	Response() http.ResponseWriter
	Request() *http.Request
}
