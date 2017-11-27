package main

import (
	"log"

	"github.com/psimoesSsimoes/query2form/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
