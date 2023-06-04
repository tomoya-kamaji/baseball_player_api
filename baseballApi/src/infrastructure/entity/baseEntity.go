package entity

import "time"

type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}
