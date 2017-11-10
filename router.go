package tuu

type Handler func(Context) error

type Router interface {
	GET(path string, ctx Context)
	POST(path string, ctx Context)
	GetRoutes() []*Route
}
