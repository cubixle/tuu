package tuu

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func New(r Router) *App {
	return &App{router: r}
}

type App struct {
	router Router
}

func (a *App) Serve() error {
	r := mux.NewRouter()

	for _, route := range a.router.GetRoutes() {
		r.Handle(route.Path, route).Methods(route.Method)
	}

	for _, route := range a.router.GetStaticRoutes() {
		r.PathPrefix(route.Path).Handler(route.Handler)
	}

	server := http.Server{
		Addr:    "",
		Handler: r,
	}

	/*ctx, cancel := sigtx.WithCancel(a.Context, syscall.SIGTERM, os.Interrupt)
	defer cancel()

	go func() {
		// gracefully shut down the application when the context is cancelled
		<-ctx.Done()
		fmt.Println("Shutting down application")

		err := a.Stop(ctx.Err())
		if err != nil {
			fmt.Println(err)
		}

		err = server.Shutdown(ctx)
		if err != nil {
			fmt.Println(err)
		}

	}()*/

	// start the web server
	return server.ListenAndServe()
}
