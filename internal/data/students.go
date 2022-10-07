//Filename: internal/data/students.go

package data

import (
	"time"
)

// holds student informatiom
// back tick character(struct tags) shows how the key should be formated
type Student struct {
	ID        string
	CreatedAt time.Time
	Name      string
	Phone     string
	Email     string
	Hobby     []string
	Color     string
	Version   int32
}
