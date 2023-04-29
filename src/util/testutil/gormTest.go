package testutil

import (
	"github.com/joho/godotenv"
	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
)

func NewTestMainDB() *gorm.DBProvider {
	err := godotenv.Load("../../../.env_test")
	if err != nil {
		panic(err)
	}
	return gorm.NewMainDB()
}
