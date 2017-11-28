package actions

import (
	"fmt"
	"net/url"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func QueryGetHandler(c buffalo.Context) error {
	values := map[string]interface{}{}
	if m, ok := c.Params().(url.Values); ok {
		for k, v := range m {
			fmt.Println(k)
			values[k] = v
		}
	}
	c.Set("values", values)

	return c.Render(200, r.HTML("query.html"))
}
