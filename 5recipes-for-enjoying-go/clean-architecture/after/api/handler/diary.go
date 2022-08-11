package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"after/api/adapter"
	"after/api/presenter"
	"after/usecase/diary"
)

func NewCreateDiaryHandler(du *diary.CreateDiaryUseCase) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Adapter
		input, err := adapter.NewCreateDiaryInputPortFromRequest(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// UseCase
		output, err := du.Execute(r.Context(), input)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		// Presenter
		if err := json.NewEncoder(w).Encode(
			presenter.NewCreateDiaryPresenter(output)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println(err)
			return
		}

	})
}
