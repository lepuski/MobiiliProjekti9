package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello valioliiga app"))

}

func (app *application) showUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display the user with an id of .. %d", id)

}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("plz register here"))
}
func (app *application) handleRegistration(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("processing registration"))
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("plz login here"))
}
func (app *application) handleLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("processing login"))
}
