//Filename: cmd/api/students.go

package main

import (
	"fmt"
	"net/http"

	"students.joelical.net/internal/data"
)

func (app *application) showMeHandler(w http.ResponseWriter, r *http.Request) {

	//create a map to hold student data
	student := data.Student{
		/*ID:            "2005113038",*/
		Name:          "Joel T. Ical",
		Phone:         "601-4113",
		Work:          "Ministry of Economic Devlopment",
		Email:         "2005113038@gmail.com",
		Hobby:         []string{"Drawing", "Reading", "Walking"},
		FavoriteColor: "Black",
		Version:       1,
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"student": student}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

// createSchoolHandler for the "GET /v1/schools/:id" endpoint
func (app *application) showStudentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	//Create a new instance of the school struct containing the ID we extracted from our URL and some sample data
	school := data.Student{
		ID:            id,
		Name:          "Joel T. Ical",
		Phone:         "601-4113",
		Work:          "Ministry of Economic Devlopment",
		Email:         "2005113038@gmail.com",
		Hobby:         []string{"Drawing", "Reading", "Walking"},
		FavoriteColor: "Black",
		Version:       1,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"school": school}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

// createSchoolHandler for the "POST /v1/schools" endpoint
func (app *application) createStudentHandler(w http.ResponseWriter, r *http.Request) {
	//our target decode destination
	var input struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
		Email string `json:"email"`
	}
	//Initalize a new json decoder Instance
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	//Display the request
	fmt.Fprintf(w, "%+v\n", input)
}
