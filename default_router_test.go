package tuu_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gitlab.com/lrodham/tuu"
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
