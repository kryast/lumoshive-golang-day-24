package model

import (
	"database/sql"
	"time"
)

type Book struct {
	ID        int
	Title     string
	Category  string
	Author    string
	Price     float64
	Discount  float64
	BookCover sql.NullString
	BookFile  sql.NullString
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
