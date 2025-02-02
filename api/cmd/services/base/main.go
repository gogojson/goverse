package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	ctx := context.Background()
	run(ctx)

}

func run(ctx context.Context) {
	// Connect to DB
	pool, err := pgxpool.New(context.Background(), os.Getenv("postgresql"))
	if err != nil {
		slog.Error(err.Error())
		panic("failed to connect to dB")
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		slog.Error("failed to ping db", "err", err)
	}

	// create mux and setup some end points

	mux := http.NewServeMux()

	// Create all end points for now
	//--------------------------------------------------------------------------
	// Admin
	mux.HandleFunc("POST /admin/add/lecture", testHandler)
	mux.HandleFunc("POST /admin/add/read/{lecture_id}", testHandler) // Add content in mark down
	mux.HandleFunc("POST /admin/add/quiz/{lecture_id}", testHandler) // Send it in JSON {problems: [{no: 1, question: "", selections: ["11","122"],answer: 3}]}
	mux.HandleFunc("POST /admin/add/code/{lecture_id}", testHandler) //Send it in JSON {problemsRepo: "", submissionRepo: ""}

	mux.HandleFunc("GET /test", testHandler)

	http.ListenAndServe("0.0.0.0:9900", mux)

}

func testHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("this is a yeah!"))

}

func PostLecture(w http.ResponseWriter, r *http.Request) {

}
