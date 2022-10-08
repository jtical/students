//Filename: internal/validator/validator.go

package validator

import (
	"net/url"
	"regexp"
)

var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	PhoneRX = regexp.MustCompile(`^\+?\(?[0-9{3}\)?\s?-\s?[0-9]{3}\s?-\s?[0-9]{4}$`)
)

// create a type that wrap our validation error map
type validator struct {
	Errors map[string]string
}

// new creates a new validator
func New() *validator {
	return &validator{
		Errors: make(map[string]string),
	}
}

// create a method to access validator type
// valid check the errors map for entries
func (v *validator) valid() bool {
	return len(v.Errors) == 0
}

// create a funct in(), which checks if an element can be found in a provided llist of elements
func In(element string, list ...string) bool {
	for i := range list {
		if element == list[i] {
			return true
		}
	}
	return false
}

// matches () returns true if a string value matches a specific regex pattern
func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

// valid website() checks if a website string value is a valid web url
func Validwesite(website string) bool {
	_, err := url.ParseRequestURI(website)
	return err == nil
}

//addError () add an error entry to Error Map

func (v *validator) AddError(key, message string) {

	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}

}

//check() performs the validation checks and call the addError
//method in turn if an error entry needs to be added

func (v *validator) Check(ok bool, key, message string) {
	if !ok {

		v.AddError(key, message)
	}

}

//Unique() check that there are no repeating values in slices

func Unique(values []string) bool {
	uniqueValues := make(map[string]bool)

	for _, value := range values {
		uniqueValues[value] = true
	}

	return len(values) == len(uniqueValues)
}
