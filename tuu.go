package tuu

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/mux"
)

type Config struct {
	IPAddr string
	Port   string
	Env    string
}

func New(r Router, cfg Config) *App {
	return &App{router: r, cfg: cfg}
}

type App struct {
	router Router
	cfg    Config
}

func (a *App) Serve() error {
	log.Printf("http server running @ http://%s:%s", a.cfg.IPAddr, a.cfg.Port)

	r := mux.NewRouter()

	for _, route := range a.router.GetRoutes() {
		r.Handle(route.Path, route).Methods(route.Method)
	}

	for _, route := range a.router.GetStaticRoutes() {
		r.PathPrefix(route.Path).Handler(route.Handler)
	}

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", a.cfg.IPAddr, a.cfg.Port),
		Handler: r,
	}

	// Check for a closing signal
	go func() {
		// Graceful shutdown
		sigquit := make(chan os.Signal, 1)
		signal.Notify(sigquit, os.Interrupt, os.Kill)

		sig := <-sigquit
		log.Printf("caught sig: %+v", sig)
		log.Printf("Gracefully shutting down server...")

		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("Unable to shut down server: %v", err)
		} else {
			log.Println("Server stopped")
		}
	}()

	// start the web server
	return server.ListenAndServe()
}
