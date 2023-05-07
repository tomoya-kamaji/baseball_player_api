package main

import (
	"fmt"
	"os"

	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
	"github.com/tomoya_kamaji/go-pkg/src/job"
	"github.com/tomoya_kamaji/go-pkg/src/util/logging"
	"github.com/urfave/cli"
)

func main() {
	baseBallJob := job.NewJobServer(gorm.NewMainDB(), logging.NewStackDriverLoggerByName("baseball-job"))
	app := &cli.App{
		Name:   "Webクローラー",
		Usage:  "This app prints 'Hello, World!'",
		Action: baseBallJob.Crawler,
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
