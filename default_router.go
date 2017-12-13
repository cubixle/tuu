package tuu

import (
	"fmt"
	"net/http"
	"strings"
)

func NewRouter() Router {
	return &DefaultRouter{}
}

type DefaultRouter struct {
	Routes       []*Route
	StaticRoutes []*StaticRoute

	prefix string
}

func (r *DefaultRouter) Prefix(path string) {
	r.prefix = path
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

func (r *DefaultRouter) addRoute(m, p string, h Handler) {
	if r.prefix != "" {
		p = fmt.Sprintf("/%s/%s", strings.Trim(r.prefix, "/"), strings.Trim(p, "/"))
	}

	r.Routes = append(r.Routes, &Route{
		Method:  m,
		Path:    p,
		Handler: h,
	})
}
