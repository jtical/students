//Filename: cmd/api/students.go

package main

import (
	"fmt"
	"net/http"
)

// createStudentHandler for the POST /v1/entries endpoint
func (app *application) createStudentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "creating a new entry..")
}

// showStudentHandler for the GET /v1/Student endpoint
func (app *application) showStudentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Student_ID: 2005113038")
	fmt.Fprintln(w, "Student_Name: Joel T. Ical")
	fmt.Fprintln(w, "Hobby: Drawing")
	fmt.Fprintln(w, "Favorite_Color: Black")
}

// createStudentHandler for the GET /v1/string endpoint
func (app *application) createRandomStringHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	//display id
	fmt.Fprintf(w, "show the details for student %d\n", id)
}
