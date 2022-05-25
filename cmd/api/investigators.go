package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pedro-git-projects/call-of-gopher/internal/data"
	"github.com/pedro-git-projects/call-of-gopher/internal/data/validator"
)

// Creates a new investigator located at /v1/investigator/<investigators name>
func (app *application) createInvestigatorsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name       string `json:"nome"`
		Age        int    `json:"idade"`
		Residence  string `json:"residencia"`
		Birthplace string `json:"nascimento"`
		Occupation string `json:"ocupacao"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	var i data.Investigator

	name := input.Name
	age := input.Age
	residence := input.Residence
	birthplace := input.Birthplace
	occupation := input.Occupation

	investigator := data.CreateInvestigator(&i, name, age, birthplace, residence, occupation)

	// performing validaton
	v := validator.New()
	if data.ValidateInvestigator(v, investigator); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/investigator/%s", investigator.Name))

	err = app.WriteJSON(w, http.StatusCreated, envelope{"investigator": investigator}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
