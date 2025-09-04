package main

import (
	"log"
	"net/http"
)


func main()  {
	const port = "8080"
	const filepathRoot = "."

	mux := http.NewServeMux()	

	server := http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	mux.HandleFunc("/healthz", handlerReadiness)

	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}