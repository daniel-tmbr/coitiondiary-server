package models

import "time"

// Todo model
type Todo struct {
	Id        int       `json:"-"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos slice, containing Todo models
type Todos []Todo
