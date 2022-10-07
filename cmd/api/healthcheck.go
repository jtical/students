//Filename: cmd/api/healthcheck.go

package main

import (
	"encoding/json"
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
	//convert our map to a json object
	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a issue and could not complete the request", http.StatusInternalServerError)
		return
	}
	//add a new line to make view on the terminal
	js = append(js, '\n')
	//specify that we will serve our response in JSON
	w.Header().Set("Content-Type", "application/json")
	//write the [] byte contain the JSON response body
	w.Write(js)

}
