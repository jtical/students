//Filename: cmd/api/healthcheck.go

package main

import (
	"fmt"
	"net/http"
)

// create a healthcheck handler post
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
