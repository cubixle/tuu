package tuu

import (
	"github.com/gorilla/mux"
)

func New(r Router) *App {
	return &App{router: r}
}

type App struct {
	router Router
}

func (a *App) Serve() {
	r := mux.NewRouter()

	for _, route := range a.router.GetRoutes() {
		r.Handle(route.Path, route).Methods(route.Method)
	}

}
