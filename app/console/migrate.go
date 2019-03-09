package console

import (
	"github.com/vzool/iris-app/app"
	model "github.com/vzool/iris-app/app/models"
	cli "gopkg.in/urfave/cli.v1"
)

// Migrate CLI
func Migrate(c *cli.Context) error {

	db := app.DB()
	defer db.Close()

	db.AutoMigrate(
		&model.City{},
		&model.Task{},
	)

	return nil
}
