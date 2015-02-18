package main

import (
	"github.com/ivpusic/neo"
)

type Person struct {
	FirstName string
	LastName  string
}

func main() {
	app := neo.App()

	app.Get("/", func(ctx *neo.Ctx) {
		person := Person{"Bob", "Roberts"}
		ctx.Res.Json(person, 200)
	})

	app.Start()
}
