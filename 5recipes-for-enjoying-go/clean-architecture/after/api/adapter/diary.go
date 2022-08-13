package adapter

import (
	"encoding/json"
	"net/http"
	"strconv"

	"after/usecase/diary"

	"github.com/gorilla/mux"
)

type createDiaryRequestBody struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewCreateDiaryInputPortFromRequest(r *http.Request) (*diary.CreateDiaryInputPort, error) {
	var input createDiaryRequestBody
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		return nil, err
	}
	return &diary.CreateDiaryInputPort{
			Title:       input.Title,
			Description: input.Description},
		nil
}

func NewGetDiaryInputPortFromRequest(r *http.Request) (*diary.GetDiaryInputPort, error) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		return nil, err
	}

	return &diary.GetDiaryInputPort{
		ID: ID,
	}, nil
}
