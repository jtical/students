//Filename: cmd/api/helpers.go

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// define a new type named envelope. empty interface means it can be any type
type envelope map[string]interface{}

// we create our method to read id
func (app *application) readIDParam(r *http.Request) (int64, error) {
	//use the "paramsfromcontect()" function to get the request context on a slice
	params := httprouter.ParamsFromContext(r.Context())
	//get the value of the id form the parameter
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

// we create our method write json to create responses
func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) error {
	//convert our map into json object
	//marshal indent to separate key form value. to appear on seprate line.
	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	//a new line to byte slices to make viewing on the terminal easier
	js = append(js, '\n')
	//add the headers
	for key, value := range headers {
		w.Header()[key] = value

	}
	// specify that we will server our responses using JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	//we write out the []byte slices containg the json response body
	w.Write(js)
	return nil
}

// dst stores decoded struct
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {

	//decode the request body into the target destinatiin
	err := json.NewDecoder(r.Body).Decode(dst)
	//check for a bad request
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshallTypeError *json.UnmarshalTypeError
		var invalidUnmarshallError *json.InvalidUnmarshalError

		//switch to check for the errors
		switch {
		//check for syntax errors
		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON(at character %d)", syntaxError.Offset)
		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")
		//Check for wrong types passed by the user
		case errors.As(err, &unmarshallTypeError):
			if unmarshallTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshallTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type(at charcater %d)", unmarshallTypeError.Offset)
		//empty body
		case errors.Is(err, io.EOF):
			return errors.New("body cannot be empty")
		//pass a non-nil pointer error
		case errors.As(err, &invalidUnmarshallError):
			panic(err)
		default:
			return err
		}
		//default

	}
	return nil
}
