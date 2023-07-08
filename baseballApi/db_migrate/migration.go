package db_migrate

import (
	"fmt"
	"log"

	"github.com/k0kubun/sqldef"
	"github.com/k0kubun/sqldef/database"
	"github.com/k0kubun/sqldef/database/mysql"
	"github.com/k0kubun/sqldef/schema"
	"github.com/urfave/cli"

	"github.com/k0kubun/sqldef/parser"
	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
)

func Migrate(cCtx *cli.Context) error {
	schemaFile := cCtx.Args().Get(0)
	// データベースへの接続情報を設定
	cfg := gorm.GetMainMySQLConfig()
	db, err := mysql.NewDatabase(database.Config{
		Host:     cfg.Host,
		Port:     cfg.Port,
		User:     cfg.User,
		Password: cfg.Password,
		DbName:   cfg.Database,
	})
	if err != nil {
		log.Fatal(err)
	}
	sqlParser := database.NewParser(parser.ParserModeMysql)
	desiredDDLs, err := sqldef.ReadFile(schemaFile)
	if err != nil {
		return fmt.Errorf("failed to read schema file: %s", err)
	}

	var dryRun bool
	if cCtx.Args().Get(1) == "apply" {
		dryRun = false
	} else {
		dryRun = true
	}
	options := &sqldef.Options{
		DesiredDDLs: desiredDDLs,
		DryRun:      dryRun,
	}

	sqldef.Run(schema.GeneratorModeMysql, db, sqlParser, options)
	return nil
}
