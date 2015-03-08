package main

import (
	"database/sql"
	"github.com/ivpusic/neo"
	_ "github.com/lib/pq"
)

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func SetupDB() *sql.DB {
	db, err := sql.Open("postgres", "dbname=lotus_todo_development sslmode=disable")
	PanicIf(err)
	return db
}

type Person struct {
	FirstName string
	LastName  string
}

type Task struct {
	Id        int
	Title     string
	Completed bool
}

func main() {
	app := neo.App()

	app.Get("/", func(ctx *neo.Ctx) {
		person := Person{"Bob", "Roberts"}
		ctx.Res.Json(person, 200)
	})

	app.Get("/tasks", func(ctx *neo.Ctx) {
		db := SetupDB()
		rows, err := db.Query("SELECT id, title, completed FROM tasks")
		PanicIf(err)
		tasks := []Task{}
		for rows.Next() {
			task := Task{}
			err := rows.Scan(&task.Id, &task.Title, &task.Completed)
			PanicIf(err)
			tasks = append(tasks, task)
		}
		defer rows.Close()
		ctx.Res.Json(tasks, 200)
	})

	app.Post("/login", func(ctx *neo.Ctx) {
		db := SetupDB()
		email, _ := ctx.Req.FormValue("email"), ctx.Req.FormValue("password")
		//TODO: add an encrypted password column to the users table
		var id, name string
		err := db.QueryRow("SELECT id, name FROM users WHERE email = $1", email).Scan(&id, &name)
		if err != nil {
			ctx.Res.Json("Not authorised", 402)
		} else {
			ctx.Res.Json("User is '"+name+"'", 200)
		}
		//TODO: Generate a token for the given user
	})

	app.Start()
}
