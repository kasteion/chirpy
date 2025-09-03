package main

import "net/http"

func main()  {
	const port = "8080"

	mux := http.NewServeMux()

	server := http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	mux.Handle("/", http.FileServer(http.Dir(".")))

	server.ListenAndServe()
}