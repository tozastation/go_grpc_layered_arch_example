package db

import "time"

// BaseModel is ...
type BaseModel struct {
	ID        uint
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-" `
}
