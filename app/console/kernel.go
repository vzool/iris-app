package console

import cli "gopkg.in/urfave/cli.v1"

// Commands is place to register all commands
func Commands() []cli.Command {

	return []cli.Command{
		{
			Name:        "run",
			Aliases:     []string{"r"},
			Action:      WebApp,
			Description: "Run the web application",
		},
		{
			Name:        "migrate",
			Aliases:     []string{"m"},
			Action:      Migrate,
			Description: "Migrate database",
		},
	}
}
