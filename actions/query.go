package actions

import (
	"fmt"
	"net/url"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func QueryHandler(c buffalo.Context) error {
	values := []string{}
	if m, ok := c.Params().(url.Values); ok {
		for _, v := range m {
			fmt.Println(v)
			values = append(values, fmt.Sprintf("%+v", v[0]))
		}
	}
	c.Set("values", values)

	return c.Render(200, r.HTML("query.html"))
}
