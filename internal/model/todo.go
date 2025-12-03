package model

import "time"

type TodoTask struct {
	ID          int       `json:"todo_id" gorm:"primary_key;auto_increment;column:todo_id"`
	UserID      int       `json:"user_id" gorm:"user_id"`
	Title       string    `json:"title" gorm:"column:title"`
	Description string    `json:"description" gorm:"column:description"`
	Completed   bool      `json:"completed" gorm:"column:completed"`
	CreatedTime time.Time `json:"due_date" gorm:"column:due_date"`
}
