package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/muhrizqiardi/linkbox/internal/model"
	"github.com/muhrizqiardi/linkbox/internal/query"
)

type LinkMediaRepository interface {
	Insert(linkID int, mediaPath string) (model.LinkMediaModel, error)
}

type linkMediaRepository struct {
	db *sqlx.DB
}

func NewLinkMediaRepository(db *sqlx.DB) *linkMediaRepository {
	return &linkMediaRepository{db}
}

func (lmr *linkMediaRepository) Insert(linkID int, mediaPath string) (model.LinkMediaModel, error) {
	stmt, err := lmr.db.Preparex(query.QueryInsertLinkMedia)
	if err != nil {
		return model.LinkMediaModel{}, err
	}

	var lm model.LinkMediaModel
	if err := stmt.Get(&lm, linkID, mediaPath); err != nil {
		return model.LinkMediaModel{}, err
	}

	return lm, nil
}
