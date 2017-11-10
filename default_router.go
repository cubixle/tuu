package tuu

import "net/http"

func NewRouter() *DefaultRouter {
	return &DefaultRouter{}
}

type DefaultRouter struct {
	Routes []*Route
}

func (r *DefaultRouter) GET(path string, h Handler) {
	r.addRoute(http.MethodGet, path, h)
}

func (r *DefaultRouter) POST(path string, h Handler) {
	r.addRoute(http.MethodPost, path, h)
}

func (r *DefaultRouter) GetRoutes() []*Route {
	return r.Routes
}

func (r *DefaultRouter) addRoute(m, p string, h Handler) {
	r.Routes = append(r.Routes, &Route{
		Method:  m,
		Path:    p,
		Handler: h,
	})
}
