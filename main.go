package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	"github.com/kasteion/chirpy/internal/database"
	_ "github.com/lib/pq"
)

func main()  {
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	port := os.Getenv("PORT")
	filepathRoot := os.Getenv("FILEPATH_ROOT")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	mux := http.NewServeMux()	

	server := http.Server{
		Addr: ":" + port,
		Handler: mux,
	}

	apiCfg := apiConfig{ 
		fileserverHits: atomic.Int32{},
		db: dbQueries,
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