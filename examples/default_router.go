package examples

import (
	"log"
	"gitlab.com/lrodham/tuu"
)

func main() {
	router := tuu.NewRouter()

	router.GET("/home", func(ctx tuu.Context) error {
		ctx.Set("template_data", "some value")

		return ctx.Render(200, render.HTML("template_name.html"))
	}) 
 
	router.POST("/login", func(ctx tuu.Context) error {
		username := ctx.Param("username")
		password := ctx.Param("password")

		log.Println(username, password)

		return nil
	})

	app := tuu.New(router)
	if err := app.Serve(); err != nil {
		panic(err)
	}
}