package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/vzool/iris-app/app"
	"github.com/vzool/iris-app/app/console"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {

	app := cli.NewApp()

	app.Name = "Iris Example App"
	app.Usage = "fight the loneliness!"
	app.Action = console.WebApp

	// Load all commands from the kernel's console
	app.Commands = console.Commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
