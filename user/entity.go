package user

import "time"

type User struct {
	ID int
	Name string
	Email string
	Photo string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}