package repositoryImpl

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/joho/godotenv"
	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
)

var (
	connectionString string
	driver           = "mysql"
	db               *sql.DB
)

func TestMain(m *testing.M) {
	initTestMain()
	defer db.Close()

	beforeEach()

	code := m.Run()

	afterEach()

	os.Exit(code)
}

func initTestMain() {
	err := godotenv.Load("../../../.env_test")
	cfg := gorm.GetMainMySQLConfig()
	if err != nil {
		panic(err)
	}
	connectionString = cfg.ConnectionString()
	db, err = sql.Open(driver, connectionString)
	if err != nil {
		panic(err)
	}
}

func beforeEach() {
	fmt.Println("before each...")
	migration()
}

func migration() {
	fmt.Println("migraion...")
	_, err := db.Exec("DROP DATABASE IF EXISTS mydb")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE DATABASE mydb")
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("goose", "-dir", "../../../migrations", driver, connectionString, "up")
	_, err = cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	fmt.Println("migration successfully")
}

func afterEach() {
	fmt.Println("after each...")
}
