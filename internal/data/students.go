//Filename: internal/data/students.go

package data

// holds student informatiom
// back tick character(struct tags) shows how the key should be formated
type Student struct {
	ID            string   `json:"id"`
	Name          string   `json:"name"`
	Phone         string   `json:"phone"`
	Work          string   `json:"work"`
	Email         string   `json:"email,omitempty"`
	Hobby         []string `json:"hobby"`
	FavoriteColor string   `json:"favorite color"`
	Version       int32    `json:"version"`
}
