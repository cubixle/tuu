package tuu_test

import (
	"net/http"
	"testing"

	"github.com/lukerodham/tuu"
	"github.com/stretchr/testify/require"
)

func Test_Route_Creation(t *testing.T) {
	r := require.New(t)

	router := tuu.NewRouter()
	router.GET("/testing", func(ctx tuu.Context) error { return nil })

	routes := router.GetRoutes()
	r.Len(routes, 1)
	route := routes[0]
	r.Equal("GET", route.Method)
	r.Equal("/testing", route.Path)
}

func Test_Static_Route_Creation(t *testing.T) {
	r := require.New(t)

	dir := "test_dir"
	router := tuu.NewRouter()
	router.Static("/test-path", http.Dir(dir))

	routes := router.GetStaticRoutes()
	r.Len(routes, 1)
	route := routes[0]
	r.Equal("/test-path", route.Path)
}
