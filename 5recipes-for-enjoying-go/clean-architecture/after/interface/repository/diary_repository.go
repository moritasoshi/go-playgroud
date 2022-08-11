package repository

import (
	"context"
	"database/sql"

	"after/domain/model"
)

type diaryRepository struct {
	db *sql.DB
}

func NewDiaryRepository(db *sql.DB) *diaryRepository {
	return &diaryRepository{db}
}

func (dr *diaryRepository) Store(ctx context.Context, diary *model.Diary) (*model.Diary, error) {
	s := `insert into diary(title, description) values ($1, $2) returning id`
	if err := dr.db.QueryRowContext(ctx, s, diary.Title, diary.Description).Scan(&diary.ID); err != nil {
		return nil, err
	}
	return diary, nil
}

func (dr *diaryRepository) FindByID(ctx context.Context, id int) (*model.Diary, error) {
	diary := &model.Diary{}
	s := `select id, title, description from diary where id = $1`
	if err := dr.db.QueryRowContext(ctx, s, id).Scan(&diary.ID, &diary.Title, &diary.Description); err != nil {
		return nil, err
	}
	return diary, nil
}
