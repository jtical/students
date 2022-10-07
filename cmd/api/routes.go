//Filename: cmd/api/routes.go

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// create a method that returns a http router
func (app *application) routes() *httprouter.Router {
	//Create a new httrouter router instance
	router := httprouter.New()
	//implement error handling in router. use our error response msg
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedesponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/random/:id", app.createRandomStringHandler)
	router.HandlerFunc(http.MethodGet, "/v1/student", app.showStudentHandler)

	return router
}
