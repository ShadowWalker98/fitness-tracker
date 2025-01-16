package main

func (app *application) routes() {
	app.mux.HandleFunc("/healthcheck", app.healthcheckHandler)
}
