package domain

import (
	"time"
)

type Book struct {
	UserID int       `bson:"userID"`
	Title  string    `bson:"title"`
	Writer string    `bson:"writer"`
	Isbn   string    `bson:"isbn"`
	Date   time.Time `bson:"date"`
}
