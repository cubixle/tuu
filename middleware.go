package tuu

type Middleware func(Handler) Handler

type MiddlewareStack struct {
	stack []Middleware
}

func (m *MiddlewareStack) Use(mw ...Middleware) {
	m.stack = append(m.stack, mw...)
}

func (m *MiddlewareStack) Get() []Middleware {
	return m.stack
}

func (m *MiddlewareStack) handler(r *Route) Handler {
	handler := r.Handler
	if len(m.stack) == 0 {
		return handler
	}

	for _, mw := range m.stack {
		handler = mw(handler)
	}

	return handler
}
