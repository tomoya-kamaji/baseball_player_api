package testutil

import (
	"database/sql"
)

func NewSqlBoilerMainDB() (*sql.DB, error) {
	return sql.Open("mysql", "root:password@localhost:3306/mydb")
}
