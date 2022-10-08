//Filename: cmd/api/students.go

package main

import (
	"net/http"

	"students.joelical.net/internal/data"
)

func (app *application) showMeHandler(w http.ResponseWriter, r *http.Request) {
	//create a map to hold student data
	student := data.Student{
		ID:            "2005113038",
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
