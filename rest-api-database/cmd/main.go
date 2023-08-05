package main

import "github.com/golanguzb70/tracing-examples/rest-api-database/app"

func main() {
	app := app.New()
	app.Run()
}
