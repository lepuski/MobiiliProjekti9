package main

import (
	"log"
	"net/http"
	"flag"
	"os"
)

//create an application struct for dependency injection
type application struct{
	errorLog 	*log.Logger
	infoLog		*log.Logger
}

func main() {

	
	//change port by specifying a flag when running main.go
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	//create a new custom logger 
	infoLog := log.New(os.Stdout, "INFO \t", log.Ldate | log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR \t", log.Ldate | log.Ltime | log.Lshortfile)

	//init a new application struct that contains the required dependencies 
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("GET /register", app.register)
	mux.HandleFunc("POST /register", app.handleRegistration)
	mux.HandleFunc("GET /snippet", app.showUser)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))


	//create a new customized http.server that uses the custom logger created above
	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Println("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
