package tuu

import (
	"net/http"

	gcontext "github.com/gorilla/context"
)

type Route struct {
	Method  string
	Path    string
	Handler Handler
}

func (r *Route) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	defer gcontext.Clear(req)

	c := NewContext(*r, res, req)

	if err := r.Handler(c); err != nil {
		c.Response().WriteHeader(500)
		c.Response().Write([]byte(err.Error()))
	}
}
