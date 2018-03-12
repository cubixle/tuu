package tuu

import (
	"net/http"

	gcontext "github.com/gorilla/context"
)

type Route struct {
	Method     string
	Path       string
	Handler    Handler
	Env        string
	Middleware MiddlewareStack
}

func (r *Route) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	defer gcontext.Clear(req)

	c := newContext(*r, res, req)

	err := r.Middleware.handler(r)(c)

	if err != nil {
		c.Response().WriteHeader(500)
		c.Response().Write([]byte(err.Error()))
	}
}

type StaticRoute struct {
	Path    string
	Handler http.Handler
}
