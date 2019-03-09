package main

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/kataras/iris"
	_ "github.com/vzool/iris-app/app"
	"github.com/vzool/iris-app/app/console"
	"github.com/vzool/iris-app/app/http"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {

	app := cli.NewApp()

	app.Name = "Iris Example App"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) error {

		app := iris.Default()

		http.Routes(app)

		// listen and serve on http://0.0.0.0:8800.
		app.Run(iris.Addr("localhost:8800"))

		return nil
	}

	// load all commands from the kernel's console
	app.Commands = console.Commands(cli.Command{
		Name:        "run",
		Aliases:     []string{"r"},
		Action:      app.Action,
		Description: "Run the web application",
	})

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
