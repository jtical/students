// Filename: cmd/api/randomize.go
package main

import (
	"crypto/rand"
	"net/http"

	"students.joelical.net/internal/data"
)

func (app *application) showRandomizeHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	rString := generateRandomString(id)
	//create a new instance of data struct
	data := data.RandomString{
		Data: rString,
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"data": data}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func generateRandomString(length int64) string {

	randomStringSource := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+_#$-!~"
	s := make([]rune, length)
	r := []rune(randomStringSource)

	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x := p.Uint64()
		y := uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)

}
