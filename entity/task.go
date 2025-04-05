package entity

import "time"

type Task struct {
	ID         int
	Title      string
	IsDone     bool
	DeadLine   time.Time
	CategoryID int
	UserID     int
}
