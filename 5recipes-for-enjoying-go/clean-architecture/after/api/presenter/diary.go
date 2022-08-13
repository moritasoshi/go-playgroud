package presenter

import (
	"after/domain/model"
	"after/usecase/diary"
)

type simpleDiaryView struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func mapDiaryToSimpleView(diary *model.Diary) *simpleDiaryView {
	return &simpleDiaryView{
		ID:          diary.ID,
		Title:       diary.Title,
		Description: diary.Description,
	}
}

type createDiaryResponse struct {
	Diary *simpleDiaryView `json:"diary"`
}

func NewCreateDiaryPresenter(output *diary.CreateDiaryOutputPort) *createDiaryResponse {
	return &createDiaryResponse{mapDiaryToSimpleView(output.Diary)}
}

func NewGetDiaryPresenter(output *diary.GetDiaryOutputPort) createDiaryResponse {
	return createDiaryResponse{mapDiaryToSimpleView(output.Diary)}
}
