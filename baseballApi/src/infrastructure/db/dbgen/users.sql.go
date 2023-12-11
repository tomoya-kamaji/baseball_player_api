// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: users.sql

package dbgen

import (
	"context"
)

const getUser = `-- name: GetUser :many
SELECT
   id, name, address, contact_number
FROM
   users
`

func (q *Queries) GetUser(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Address,
			&i.ContactNumber,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}