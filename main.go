package main

import "github.com/babell00/gamedealbackend/app"

func main() {
	app := app.App{}
	app.Initialize()

	app.Run("8080")
}
