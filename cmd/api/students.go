//Filename: cmd/api/students.go

package main

import (
	"fmt"
	"net/http"
	"time"

	"students.joelical.net/internal/data"
)

// createStudentHandler for the POST /v1/student endpoint
func (app *application) createStudentHandler(w http.ResponseWriter, r *http.Request) {
	//our target decode destination
	var input struct {
		Name  string   `json:"name"`
		Email string   `json:"email"`
		Phone string   `json:"phone"`
		Hobby []string `json:"hobby"`
	}
	//initialize a new json.decoder instance
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	//Display the request
	fmt.Fprintf(w, "%+v\n", input)

}

// showStudentHandler for the GET /v1/Student endpoint
func (app *application) showStudentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	//Create a new instance of the student struct containing the ID we extracted from our URL and some sample data
	student := data.Student{
		ID:        id,
		CreatedAt: time.Now(),
		Name:      "Joel T. Ical",
		Phone:     "601-4113",
		Email:     "2005113038@gmail.com",
		Hobby:     []string{"Drawing", "Reading"},
		Color:     "black",
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"student": student}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

// createStudentHandler for the GET /v1/string endpoint
func (app *application) createRandomStringHandler(w http.ResponseWriter, r *http.Request) {
	/*id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	student := data.Student{
		ID:        id,
		CreatedAt: time.Now(),
		Name:      "Joel T. Ical",
		Phone:     "601-4113",
		Email:     "2005113038@gmail.com",
		Hobby:     []string{"Drawing", "Reading"},
		Color:     "black",
		Version:   1,
	}
	err = app.writeJSON(w, http.StatusOK, student, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "the server encoutered a proble", http.StatusInternalServerError)
	}*/

}
