package tuu

import (
	"net/http"

	"github.com/gorilla/mux"

	gcontext "github.com/gorilla/context"
)

type Route struct {
	Method     string
	Path       string
	Handler    Handler
	MuxHandler mux.Route
	Env        string
}

func (r *Route) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	defer gcontext.Clear(req)

	c := NewContext(*r, res, req, r.Env)

	if err := r.Handler(c); err != nil {
		c.Response().WriteHeader(500)
		c.Response().Write([]byte(err.Error()))
	}
}

type StaticRoute struct {
	Path    string
	Handler http.Handler
}
