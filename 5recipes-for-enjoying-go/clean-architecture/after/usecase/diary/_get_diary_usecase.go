// package diary
//
// import (
// 	"context"
//
// 	"after/domain/model"
// 	"after/usecase/repository"
// )
//
// type GetDiaryInputPort struct {
// 	ID int
// }
//
// type GetDiaryOutputPort struct {
// 	Diary *model.Diary
// }
//
// type GetDiaryUseCase struct {
// 	diaryRepo repository.DiaryRepository
// }
//
// func NewGetDiaryUseCase(dr repository.DiaryRepository) *GetDiaryUseCase {
// 	return &GetDiaryUseCase{dr}
// }
//
// func (du *GetDiaryUseCase) Execute(
// 	ctx context.Context,
// 	in *GetDiaryInputPort) (*GetDiaryOutputPort, error) {
// 	diary, err := du.diaryRepo.FindByID(ctx, in.ID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &GetDiaryOutputPort{diary}, nil
// }
