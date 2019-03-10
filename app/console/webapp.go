package console

import (
	"github.com/kataras/iris"
	"github.com/vzool/iris-app/app/http"
	cli "gopkg.in/urfave/cli.v1"
)

// WebApp Command
func WebApp(c *cli.Context) error {

	app := iris.Default()

	http.Routes(app)

	// listen and serve on http://0.0.0.0:8800.
	app.Run(
		iris.Addr("localhost:8800"),
		// iris.WithoutBanner,
		// iris.WithoutServerError(iris.ErrServerClosed),
	)

	return nil
}
