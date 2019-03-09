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

var Env map[string]string

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

	app.Commands = []cli.Command{
		{
			Name:    "migrate",
			Aliases: []string{"m"},
			Action:  console.Migrate,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
