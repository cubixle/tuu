package examples

import (
	"log"

	"github.com/gobuffalo/buffalo/render"
	"github.com/lukerodham/tuu"
)

func main() {
	// The render we currently use is the one built by gobuffalo.
	r := render.New(render.Options{})

	router := tuu.NewRouter()
	router.SetEnv("dev")
	router.GET("/home", func(ctx tuu.Context) error {
		ctx.Set("template_data", "some value")

		return ctx.Render(200, r.HTML("template_name.html"))
	})

	router.POST("/login", func(ctx tuu.Context) error {
		username := ctx.Param("username")
		password := ctx.Param("password")

		log.Println(username, password)

		return nil
	})

	app := tuu.New(router)
	err := app.Serve(tuu.Config{
		IPAddr: "127.0.0.1",
		Port:   "8080",
		Env:    "dev",
	}
	
	if err != nil {
		panic(err)
	}
}
