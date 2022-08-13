package diary

import (
	"context"

	"after/domain/model"
	"after/usecase/repository"
)

type (
	// UseCase(ビジネスロジック)の引数
	GetDiaryInputPort struct {
		ID int
	}
	// UseCase(ビジネスロジック)の戻り値
	GetDiaryOutputPort struct {
		Diary *model.Diary
	}
)

// 記事を取得するユースケース
type GetDiaryUseCase struct {
	diaryRepo repository.DiaryRepository
}

// TODO: initにしたい
func NewGetDiaryUseCase(dr repository.DiaryRepository) *GetDiaryUseCase {
	return &GetDiaryUseCase{dr}
}

// 記事を取得するユースケース
func (du *GetDiaryUseCase) Execute(ctx context.Context, in *GetDiaryInputPort) (*GetDiaryOutputPort, error) {
	diary, err := du.diaryRepo.FindByID(ctx, in.ID)
	if err != nil {
		return nil, err
	}
	return &GetDiaryOutputPort{diary}, nil
}
