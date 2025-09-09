package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

func main()  {
	const port = "8080"
	const filepathRoot = "."

	mux := http.NewServeMux()	

	server := http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	apiCfg := apiConfig{ 
		fileserverHits: atomic.Int32{},
	}

	mux.HandleFunc("GET /api/healthz", handlerReadiness)

	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)

	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)

	mux.HandleFunc("POST /api/validate_chirp", handlerValidateChirp)

	appHandler := http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(appHandler))

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(server.ListenAndServe())
}