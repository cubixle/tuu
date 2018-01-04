package tuu

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gobuffalo/buffalo/render"
)

type Context interface {
	context.Context
	Response() http.ResponseWriter
	Request() *http.Request
	Params() url.Values
	Param(key string) string
	Set(key string, value interface{})
	Render(status int, rr render.Renderer) error
	Redirect(status int, url string) error
	Env() string
}
