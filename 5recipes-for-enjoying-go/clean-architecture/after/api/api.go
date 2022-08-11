package api

import (
	"database/sql"

	"github.com/gorilla/mux"

	"after/api/handler"
	"after/interface/repository"
	"after/usecase/diary"
)

func BuildRouter(db *sql.DB) *mux.Router {
	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	buildProjectRoutes(r, db)
	return r
}

func buildProjectRoutes(r *mux.Router, db *sql.DB) {
	dr := repository.NewDiaryRepository(db)

	r.Handle("/diary", handler.NewCreateDiaryHandler(diary.NewCreateDiaryUseCase(dr))).Methods("POST")

}
