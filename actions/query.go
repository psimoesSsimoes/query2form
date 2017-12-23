package actions

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func QueryGetHandler(c buffalo.Context) error {

	var (
		status  string
		problem bool
	)
	values := map[string]int{}

	if m, ok := c.Params().(url.Values); ok {
		for k, v := range m {
			fmt.Println(v)
			n, err := strconv.Atoi(v[0])
			if err == nil {
				values[k] = n
			}
		}
	}

	switch {
	case values["measured"] < values["expected"]:
		status = fmt.Sprintf("Problem. Produced %.2f percent less than was planned", float32(values["measured"]*100)/float32(values["expected"]))
		problem = true
	case values["measured"] == values["expected"]:
		status = "Ok. Production corresponds with what was expected"
		problem = false
	case values["measured"] > values["expected"]:
		problem = true

		status = fmt.Sprintf("Problem. Produced %.2f percent more than was planned", float32((values["measured"]-values["expected"])*100)/float32(values["expected"]))

	}

	t := time.Now()
	timestamp := fmt.Sprintf(t.Format("2006-01-02 15:04:05"))

	c.Set("timestamp", timestamp)
	c.Set("status", status)
	c.Set("problem", problem)
	c.Set("measured", values["measured"])
	c.Set("expected", values["expected"])
	return c.Render(200, r.HTML("query.html"))
}
