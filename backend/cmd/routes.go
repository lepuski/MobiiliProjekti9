package main

import "net/http"

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)
	

	mux.HandleFunc("GET /user/register", app.userRegister)
	mux.HandleFunc("POST /user/register", app.userRegisterPost)
	mux.HandleFunc("GET /user/login", app.userLogin)
	mux.HandleFunc("POST /user/login", app.userLoginPost)
	mux.HandleFunc("POST /user/logout", app.userLogoutPost)

	mux.HandleFunc("GET /user/view/{id}", app.showUser)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return app.sessionManager.LoadAndSave(app.recoverPanic(app.logRequest(commonHeaders(mux))))
}
