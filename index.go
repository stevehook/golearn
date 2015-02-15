package main

import (
	"github.com/ivpusic/neo"
)

func main() {
	app := neo.App()

	app.Get("/", func(ctx *neo.Ctx) {
		ctx.Res.Text("I am Neo Programmer", 200)
	})

	app.Start()
}
