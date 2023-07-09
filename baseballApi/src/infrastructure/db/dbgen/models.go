// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0

package dbgen

import (
	"database/sql"
)

type Book struct {
	ID              int32
	Title           sql.NullString
	Author          sql.NullString
	Publisher       sql.NullString
	PublicationYear sql.NullInt32
	Genre           sql.NullString
}

type Loan struct {
	ID         int32
	BookID     sql.NullInt32
	UserID     sql.NullInt32
	LoanDate   sql.NullTime
	DueDate    sql.NullTime
	ReturnDate sql.NullTime
}

type User struct {
	ID            int32
	Name          sql.NullString
	Address       sql.NullString
	ContactNumber sql.NullString
}