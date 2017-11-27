package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/psimoesSsimoes/query2form/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
