package console

import cli "gopkg.in/urfave/cli.v1"

// Commands is place to register all commands
func Commands(defaultCommand cli.Command) []cli.Command {

	return []cli.Command{
		defaultCommand,
		{
			Name:        "migrate",
			Aliases:     []string{"m"},
			Action:      Migrate,
			Description: "Migrate database",
		},
	}
}
