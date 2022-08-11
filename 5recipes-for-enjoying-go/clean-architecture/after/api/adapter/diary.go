package adapter

import (
	"encoding/json"
	"net/http"

	"after/usecase/diary"
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
