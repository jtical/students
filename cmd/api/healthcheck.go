//Filename: cmd/api/healthcheck.go

package main

import (
	"net/http"
)

// create a healthcheck handler post
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//create a map to hold healthcheck data
	//we can make it an enevlope as its a map
	data := map[string]string{
		"status":      "available",
		"enviornment": app.config.env,
		"version":     version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Panicln(err)
		http.Error(w, "The server encounter a problem and could not process request", http.StatusInternalServerError)
		return
	}

}
