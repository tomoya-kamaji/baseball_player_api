package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/db/schema"
	"github.com/urfave/cli"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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
			Action: schema.Migrate,
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
