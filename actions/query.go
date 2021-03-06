package actions

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/psimoesSsimoes/query2form/models"
)

type Report struct {
	id        int    `form:"id"`
	expected  int    `form:"expected"`
	measured  int    `form:"measured"`
	eventcode int    `form:"eventcode"`
	notes     string `form:"notes"`
	problem   string `form:"problem"`
	submit    string `form:"submit"`
}

var report models.Report

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
			n, err := strconv.Atoi(v[0])
			if err == nil {
				values[k] = n
			}
		}
	}

	switch {
	case values["measured"] < values["expected"]:
		status = fmt.Sprintf("Problem. Produced %.2f percent less than was planned", 100-float32(values["measured"]*100)/float32(values["expected"]))
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

	report = models.Report{}
	report.FillReportInt(values["expected"], values["measured"], values["eventcode"], values["id"], timestamp)

	c.Set("timestamp", timestamp)
	c.Set("status", status)
	c.Set("problem", problem)
	c.Set("id", values["id"])
	c.Set("eventcode", values["eventcode"])
	c.Set("measured", values["measured"])
	c.Set("expected", values["expected"])

	return c.Render(200, r.HTML("query.html"))
}

func QueryPostHandler(c buffalo.Context) error {

	var (
		notes   string
		problem string
	)

	if val, ok := c.Request().Form["notes"]; ok {

		notes = val[0]
	}

	if val, ok := c.Request().Form["problem"]; ok {

		problem = val[0]
	}

	report.FillReportString(notes, problem)
	fmt.Println(report)
	models.ReportCSV(report)
	c.Flash().Add("success", "Report was successfully created!")
	return c.Render(200, r.HTML("saved.html"))

}
