package main

import (
	"FitnessTracker/internal/data"
	"FitnessTracker/internal/validator"
	"fmt"
	"net/http"
)

func (app *application) registerUserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(405)
		_, err := fmt.Fprintf(w, "GET Method not allowed")
		if err != nil {
			app.logger.Println("Error occurred: Incorrect method type while registering user")
		}
	case "POST":
		input := struct {
			FirstName string `json:"first-name"`
			LastName  string `json:"last-name"`
			Email     string `json:"email"`
			Password  string `json:"password"`
		}{}
		err := app.readJSON(w, r, &input)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		// copying the values into the user struct
		user := data.User{
			ID:        0,
			FirstName: input.FirstName,
			LastName:  input.LastName,
			Email:     input.Email,
			Password:  input.Password,
			Salt:      "",
		}
		// initialising the validation object
		v := validator.New()
		// validating the user struct
		data.ValidateUser(v, user)
		if !v.Valid() {
			app.failedValidationResponse(w, r, v.Errors)
			return
		}
		// adding user if everything looks good
		err = app.models.UserModel.Insert(&user)
		if err != nil {
			http.Error(w, "Error occurred while inserting user", http.StatusInternalServerError)
			return
		}

		_, err = fmt.Fprintf(w, "User added successfully!")
		if err != nil {
			app.logger.Println("Error occurred while adding user")
		}
	}
}
