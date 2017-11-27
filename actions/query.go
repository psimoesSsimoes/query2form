package actions

import (
	"fmt"
	"reflect"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func QueryHandler(c buffalo.Context) error {
	x := reflect.ValueOf(c.Params()).MapKeys()
	values := []string{}
	for _, v := range x {
		fmt.Println("%T", v)
		values = append(values, fmt.Sprintf("%+v", v))
	}
	c.Set("values", values)

	return c.Render(200, r.HTML("query.html"))
}
