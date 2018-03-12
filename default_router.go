package tuu

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

func NewRouter(opts ...RouterOption) Router {
	router := &DefaultRouter{}
	router.Options = &RouterOptions{}
	for _, opt := range opts {
		opt(router.Options)
	}

	return router
}

type DefaultRouter struct {
	Routes       []*Route
	StaticRoutes []*StaticRoute
	Options      *RouterOptions
}

func (r *DefaultRouter) GET(path string, h Handler) {
	r.addRoute(http.MethodGet, path, h)
}

func (r *DefaultRouter) POST(path string, h Handler) {
	r.addRoute(http.MethodPost, path, h)
}

func (r *DefaultRouter) Static(path string, root http.FileSystem) {
	r.StaticRoutes = append(r.StaticRoutes, &StaticRoute{
		Path:    path,
		Handler: http.StripPrefix(path, http.FileServer(root)),
	})
}

func (r *DefaultRouter) NotFound(path string, h Handler) {

}

func (r *DefaultRouter) GetRoutes() []*Route {
	return r.Routes
}

func (r *DefaultRouter) GetStaticRoutes() []*StaticRoute {
	return r.StaticRoutes
}

func (r *DefaultRouter) GetOptions() *RouterOptions {
	return r.Options
}

func (r *DefaultRouter) addRoute(m, p string, h Handler) {
	if r.Options.Prefix != "" {
		p = fmt.Sprintf("/%s/%s", strings.Trim(r.Options.Prefix, "/"), strings.Trim(p, "/"))
	}

	r.Routes = append(r.Routes, &Route{
		Method:     m,
		Path:       p,
		Handler:    h,
		Env:        r.Options.Env,
		Middleware: r.Options.MiddlewareStack,
		Logger:     r.Options.Logger,
	})
}

type RouterOptions struct {
	Env             string
	Prefix          string
	MiddlewareStack MiddlewareStack
	Logger          *logrus.Logger
}

type RouterOption func(*RouterOptions)

func RouterEnv(env string) RouterOption {
	return func(o *RouterOptions) {
		o.Env = env
	}
}

func RouterPrefix(p string) RouterOption {
	return func(o *RouterOptions) {
		o.Prefix = p
	}
}

func RouterMiddleware(ms MiddlewareStack) RouterOption {
	return func(o *RouterOptions) {
		o.MiddlewareStack = ms
	}
}

func RouterLogger(l *logrus.Logger) RouterOption {
	return func(o *RouterOptions) {
		o.Logger = l
	}
}
