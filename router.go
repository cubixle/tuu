package tuu

import "net/http"

type Handler func(Context) error

type Router interface {
	GET(path string, h Handler)
	POST(path string, h Handler)
	Static(path string, root http.FileSystem)

	GetRoutes() []*Route
	GetStaticRoutes() []*StaticRoute
}
