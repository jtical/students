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
	js := `{"Student_ID": "2005113038", "Student_Name": "Joel T. Ical", "Hobby": "Drawing", "Favorite_Color": "Black"}`
	js = fmt.Sprintf(js)
	//specify that we will serve our response in JSON
	w.Header().Set("Content-Type", "application/json")
	//write the JSON as the HTTP response body
	w.Write([]byte(js))
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
