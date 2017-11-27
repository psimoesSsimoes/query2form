package actions

import (
	"fmt"
	"reflect"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/render"
)

// HomeHandler is a default handler to serve up
// a home page.
func QueryHandler(c buffalo.Context) error {
	x := reflect.ValueOf(c.Params()).MapKeys()

	for k, v := range x {
		fmt.Println("%s %s", k, v)
	}

	return c.Render(200, render.String(fmt.Sprintf("%v\n", c.Params())))
}
