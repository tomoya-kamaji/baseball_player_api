package main

import (
	"fmt"
	"os"

	"github.com/tomoya_kamaji/go-pkg/db_migrate"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:  "greet",
			Usage: "fight the loneliness!",
			Action: func(*cli.Context) error {
				fmt.Println("Hello friend!")
				return nil
			},
		},
		{
			Name:   "migration",
			Usage:  "DBマイグレーション",
			Action: db_migrate.Migrate,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
