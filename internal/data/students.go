//Filename: internal/data/students.go

package data

import (
	"time"
)

// holds student informatiom
// back tick character(struct tags) shows how the key should be formated
type Student struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Email     string    `json:"email,omitempty"`
	Hobby     []string  `json:"hobby"`
	Color     string    `json:"color"`
	Version   int32     `json:"version"`
}
