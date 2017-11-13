package tuu

import "net/http"

type Handler func(Context) error

type Router interface {
	Prefix(path string)
	GET(path string, h Handler)
	POST(path string, h Handler)
	Static(path string, root http.FileSystem)
	NotFound(path string, h Handler)

	GetRoutes() []*Route
	GetStaticRoutes() []*StaticRoute
}
