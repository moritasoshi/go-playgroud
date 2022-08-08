package main

import (
	"before/config"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Diary struct {
	ID          int    `db:"id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}

type DiaryRequest struct {
	Title       string `json:"title"`
	Description string `db:"description"`
}

func NewDB() (*sql.DB, error) {
	c := config.Config
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBPassword,
		c.DBName,
	)
	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func NewServer(handler http.Handler) *http.Server {
	c := config.Config
	return &http.Server{
		Handler:      handler,
		Addr:         fmt.Sprintf("0.0.0.0:%s", c.ServerPort),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func addDiary(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input DiaryRequest
		fmt.Printf("input: %v\n", input)
		fmt.Printf("&input: %v\n", &input)
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, err := db.Exec(
			`insert into diary(title, description) values ($1,$2)`, input.Title, input.Description,
		)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	})
}

func main() {
	conn, err := NewDB()
	if err != nil {
		log.Fatalf("Failed to initialize the database: %v", err)
	}
	defer conn.Close()

	r := mux.NewRouter().PathPrefix("/api").Subrouter()

	r.Handle("/diary", addDiary(conn)).Methods("POST")

	srv := NewServer(r)

	log.Printf("Serving on localhost:%v\n", config.Config.ServerPort)
	log.Fatal(srv.ListenAndServe())
}
