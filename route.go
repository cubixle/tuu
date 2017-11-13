package tuu

import (
	"github.com/gorilla/mux"
	"net/http"

	gcontext "github.com/gorilla/context"
)

type Route struct {
	Method  string
	Path    string
	Handler Handler
	MuxHandler mux.Route
}

func (r *Route) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	defer gcontext.Clear(req)

	c := NewContext(*r, res, req)

	if err := r.Handler(c); err != nil {
		c.Response().WriteHeader(500)
		c.Response().Write([]byte(err.Error()))
	}
}

type StaticRoute struct {
	Path    string
	Handler http.Handler
}
