package main

func (app *application) routes() {
	app.mux.HandleFunc("/healthcheck", app.healthcheckHandler)
	app.mux.HandleFunc("/add-user", app.registerUserHandler)
}
