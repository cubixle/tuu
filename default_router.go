package tuu

import "net/http"

func NewRouter() *DefaultRouter {
	return &DefaultRouter{}
}

type DefaultRouter struct {
	Routes       []*Route
	StaticRoutes []*StaticRoute
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

func (r *DefaultRouter) GetRoutes() []*Route {
	return r.Routes
}

func (r *DefaultRouter) GetStaticRoutes() []*StaticRoute {
	return r.StaticRoutes
}

func (r *DefaultRouter) addRoute(m, p string, h Handler) {
	r.Routes = append(r.Routes, &Route{
		Method:  m,
		Path:    p,
		Handler: h,
	})
}
