package model

import "time"

type Book struct {
	ID        int
	Title     string
	Category  string
	Author    string
	Price     float64
	Discount  float64
	BookCover string
	BookFile  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
