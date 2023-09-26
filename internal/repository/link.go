package repository

import (
	"errors"
	"strconv"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/jmoiron/sqlx"
	"github.com/muhrizqiardi/linkbox/internal/model"
	"github.com/muhrizqiardi/linkbox/internal/query"
	// "github.com/muhrizqiardi/linkbox/pkg/common"
)

const (
	OrderByCreatedAt = "created_at"
	OrderByUpdatedAt = "updated_at"

	SortASC  = "asc"
	SortDESC = "desc"
)

type LinkRepository interface {
	CreateLink(
		url string,
		title string,
		description string,
		user_id int,
		folder_id int,
	) (model.LinkModel, error)
	GetOneLinkByID(id int) (model.LinkModel, error)
	SearchFullText(userID int, query string) ([]model.LinkModel, error)
	GetManyLinksInsideDefaultFolder(
		userID int,
		limit int,
		offset int,
		orderBy string,
		sort string,
	) ([]model.LinkModel, error)
	GetManyLinksInsideFolder(
		userID int,
		folder_id int,
		limit int,
		offset int,
		orderBy string,
		sort string,
	) ([]model.LinkModel, error)
	UpdateOneLinkByID(
		id int,
		url string,
		title string,
		description string,
		user_id int,
		folder_id int,
	) (model.LinkModel, error)
	DeleteOneLinkByID(id int) (model.LinkModel, error)
}

type linkRepository struct {
	db  *sqlx.DB
	rsc *redisearch.Client
}

func NewLinkRepository(db *sqlx.DB, rsc *redisearch.Client) *linkRepository {
	return &linkRepository{db, rsc}
}

func (r *linkRepository) CreateLink(
	url string,
	title string,
	description string,
	user_id int,
	folder_id int,
) (model.LinkModel, error) {
	stmt, err := r.db.Preparex(query.QueryCreateLink)
	if err != nil {
		return model.LinkModel{}, err
	}

	var link model.LinkModel
	if err := stmt.Get(&link, url, title, description, user_id, folder_id); err != nil {
		return model.LinkModel{}, err
	}

	return link, nil
}

func (r *linkRepository) SearchFullText(userID int, query string) ([]model.LinkModel, error) {
	docs, total, err := r.rsc.Search(
		redisearch.
			NewQuery(query).
			SetInFields("title", "description", "url").
			SetReturnFields(
				"id",
				"url",
				"title",
				"description",
				"user_id",
				"folder_id",
				"created_at",
				"updated_at",
			),
	)
	if err != nil {
		return []model.LinkModel{}, err
	}

	ll := make([]model.LinkModel, 0, total)
	for _, e := range docs {
		idStr, ok := e.Properties["id"].(string)
		if !ok {
			return []model.LinkModel{}, errors.New("field id does not exist in search result")
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return []model.LinkModel{}, err
		}
		url, ok := e.Properties["url"].(string)
		if !ok {
			return []model.LinkModel{}, errors.New("field url does not exist in search result")
		}
		title, ok := e.Properties["title"].(string)
		if !ok {
			return []model.LinkModel{}, errors.New("field title does not exist in search result")
		}
		description, ok := e.Properties["description"].(string)
		if !ok {
			return []model.LinkModel{}, errors.New("field description does not exist in search result")
		}
		userIDStr, ok := e.Properties["user_id"].(string)
		if !ok {
			return []model.LinkModel{}, errors.New("field user_id does not exist in search result")
		}
		resUserID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return []model.LinkModel{}, err
		}
		if resUserID != userID {
			continue
		}
		folderIDStr, ok := e.Properties["folder_id"].(string)
		if !ok {
			return []model.LinkModel{}, errors.New("field folder_id does not exist in search result")
		}
		folderID, err := strconv.Atoi(folderIDStr)
		if err != nil {
			return []model.LinkModel{}, err
		}
		ll = append(ll, model.LinkModel{
			ID:          id,
			URL:         url,
			Title:       title,
			Description: description,
			UserID:      resUserID,
			FolderID:    folderID,
		})
	}

	return ll, nil
}

func (r *linkRepository) GetOneLinkByID(id int) (model.LinkModel, error) {
	stmt, err := r.db.Preparex(query.QueryGetOneLinkByID)
	if err != nil {
		return model.LinkModel{}, err
	}

	var link model.LinkModel
	if err := stmt.Get(&link, id); err != nil {
		return model.LinkModel{}, err
	}

	return link, nil
}

func (r *linkRepository) GetManyLinksInsideDefaultFolder(
	userID int,
	limit int,
	offset int,
	orderBy string,
	sort string,
) ([]model.LinkModel, error) {
	var q string = query.QueryGetManyLinksInsideDefaultFolder_OrderByUpdatedAtSortDESC
	switch orderBy {
	case OrderByCreatedAt:
		switch sort {
		case SortASC:
			q = query.QueryGetManyLinksInsideDefaultFolder_OrderByCreatedAtSortASC
			break
		case SortDESC:
			q = query.QueryGetManyLinksInsideDefaultFolder_OrderByCreatedAtSortDESC
			break
		}
		break
	case OrderByUpdatedAt:
		switch sort {
		case SortASC:
			q = query.QueryGetManyLinksInsideDefaultFolder_OrderByUpdatedAtSortASC
			break
		case SortDESC:
			q = query.QueryGetManyLinksInsideDefaultFolder_OrderByUpdatedAtSortDESC
			break
		}
		break
	}

	stmt, err := r.db.Preparex(q)
	if err != nil {
		return []model.LinkModel{}, err
	}

	var link []model.LinkModel
	if err := stmt.Select(&link, userID, limit, offset); err != nil {
		return []model.LinkModel{}, err
	}

	return link, nil
}

func (r *linkRepository) GetManyLinksInsideFolder(
	userID int,
	folderId int,
	limit int,
	offset int,
	orderBy string,
	sort string,
) ([]model.LinkModel, error) {
	var q string
	switch orderBy {
	case OrderByCreatedAt:
		switch sort {
		case SortASC:
			q = query.QueryGetManyLinksInsideFolder_OrderByCreatedAtSortASC
			break
		case SortDESC:
			q = query.QueryGetManyLinksInsideFolder_OrderByCreatedAtSortDESC
			break
		}
		break
	case OrderByUpdatedAt:
		switch sort {
		case SortASC:
			q = query.QueryGetManyLinksInsideFolder_OrderByUpdatedAtSortASC
			break
		case SortDESC:
			q = query.QueryGetManyLinksInsideFolder_OrderByUpdatedAtSortDESC
			break
		}
		break
	}

	stmt, err := r.db.Preparex(q)
	if err != nil {
		return []model.LinkModel{}, err
	}

	var link []model.LinkModel
	if err := stmt.Select(&link, userID, folderId, limit, offset); err != nil {
		return []model.LinkModel{}, err
	}

	return link, nil
}

func (r *linkRepository) UpdateOneLinkByID(
	id int,
	url string,
	title string,
	description string,
	user_id int,
	folder_id int,
) (model.LinkModel, error) {
	stmt, err := r.db.Preparex(query.QueryUpdateOneLinkByID)
	if err != nil {
		return model.LinkModel{}, err
	}

	var link model.LinkModel
	if err := stmt.Get(&link, id, url, title, description, user_id, folder_id); err != nil {
		return model.LinkModel{}, err
	}

	return link, nil
}

func (r *linkRepository) DeleteOneLinkByID(id int) (model.LinkModel, error) {
	stmt, err := r.db.Preparex(query.QueryDeleteOneLinkByID)
	if err != nil {
		return model.LinkModel{}, err
	}

	var link model.LinkModel
	if err := stmt.Get(&link, id); err != nil {
		return model.LinkModel{}, err
	}

	return link, nil
}
