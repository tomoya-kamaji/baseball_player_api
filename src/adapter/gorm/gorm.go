package gorm

import (
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v4"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBProvider struct {
	db *gorm.DB
}

var mainDB *DBProvider

// NewMainDB ... connect MySQL
func NewMainDB() *DBProvider {
	if mainDB != nil {
		return mainDB
	}
	cfg := GetMainMySQLConfig()
	d := mySQL(*cfg)
	mainDB = &DBProvider{db: d}
	return mainDB
}

func mySQL(cfg MainMySQLConfig) *gorm.DB {
	dURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database)

	var (
		gormDB        *gorm.DB
		err           error
		retryStrategy = backoff.NewExponentialBackOff()
	)

	retryStrategy.MaxElapsedTime = time.Minute
	if err := backoff.Retry(func() error {
		gormDB, err = gorm.Open(mysql.Open(dURL+"?parseTime=true&loc=Asia%2FTokyo"), &gorm.Config{})
		if err != nil {
			return err
		}
		return nil
	}, retryStrategy); err != nil {
		panic(err)
	}

	if cfg.DebugMode {
		gormDB = gormDB.Debug()
	}
	gormDB = gormDB.Set("gorm:save_associations", false) // NOTE: デフォルトでは関連テーブルの保存を許可しない

	return gormDB
}
