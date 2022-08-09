package main

import (
	"before/config"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

func handleError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusBadRequest)
}

func addDiary(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var input DiaryRequest
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

func editDiary(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var input DiaryRequest
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = db.Exec(`update diary set title = $1, description = $2 where id = $3`,
			input.Title, input.Description, ID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	})
}

func deleteDiary(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		handleError(w, err)
		_, err = db.Exec(`delete from diary where id = $1`, ID)
		handleError(w, err)
	})
}

func getDiary(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ID, err := strconv.Atoi(vars["id"])
		handleError(w, err)

		var diary Diary

		rows := db.QueryRow(`select id, title, description from diary where id = $1`, ID)
		if err := rows.Scan(
			&diary.ID, &diary.Title, &diary.Description,
		); err != nil {
			handleError(w, err)
		}
		fmt.Fprint(w, diary)
	})
}

func getDiaries(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`select id, title, description from diary`)
		if err != nil {
			handleError(w, err)
			return
		}

		defer rows.Close()

		var diaries []Diary
		for rows.Next() {
			var diary Diary
			if err := rows.Scan(&diary.ID, &diary.Title, &diary.Description); err != nil {
				handleError(w, err)
				return
			}
			diaries = append(diaries, diary)
		}

		json, err := json.Marshal(diaries)
		if err != nil {
			handleError(w, err)
			return
		}

		fmt.Fprint(w, string(json))

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
	r.Handle("/diary/{id}", getDiary(conn)).Methods("GET")
	r.Handle("/diary/{id}", editDiary(conn)).Methods("PUT")
	r.Handle("/diary/{id}", deleteDiary(conn)).Methods("DELETE")
	r.Handle("/diaries", getDiaries(conn)).Methods("GET")

	srv := NewServer(r)

	log.Printf("Serving on localhost:%v\n", config.Config.ServerPort)
	log.Fatal(srv.ListenAndServe())
}
