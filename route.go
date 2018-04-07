package tuu

import (
	"net/http"

	"github.com/gorilla/sessions"

	gcontext "github.com/gorilla/context"
	"github.com/sirupsen/logrus"
)

type Route struct {
	Method      string
	Path        string
	Handler     Handler
	Env         string
	Middleware  MiddlewareStack
	Logger      *logrus.Logger
	Session     sessions.Store
	SessionName string
}

func (r *Route) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	defer gcontext.Clear(req)

	c := newContext(*r, res, req)
	defer c.Flash().persist(c.Session())

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
