package link

import (
	"errors"
	"strconv"

	"github.com/RediSearch/redisearch-go/redisearch"
	"github.com/jmoiron/sqlx"
	"github.com/muhrizqiardi/linkbox/pkg/common"
)

const (
	OrderByCreatedAt = "created_at"
	OrderByUpdatedAt = "updated_at"

	SortASC  = "asc"
	SortDESC = "desc"
)

type repository struct {
	db  *sqlx.DB
	rsc *redisearch.Client
}

func NewRepository(db *sqlx.DB, rsc *redisearch.Client) *repository {
	return &repository{db, rsc}
}

func (r *repository) CreateLink(
	url string,
	title string,
	description string,
	user_id int,
	folder_id int,
) (common.LinkEntity, error) {
	stmt, err := r.db.Preparex(QueryCreateLink)
	if err != nil {
		return common.LinkEntity{}, err
	}

	var link common.LinkEntity
	if err := stmt.Get(&link, url, title, description, user_id, folder_id); err != nil {
		return common.LinkEntity{}, err
	}

	return link, nil
}

func (r *repository) SearchFullText(userID int, query string) ([]common.LinkEntity, error) {
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
		return []common.LinkEntity{}, err
	}

	ll := make([]common.LinkEntity, 0, total)
	for _, e := range docs {
		idStr, ok := e.Properties["id"].(string)
		if !ok {
			return []common.LinkEntity{}, errors.New("field id does not exist in search result")
		}
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return []common.LinkEntity{}, err
		}
		url, ok := e.Properties["url"].(string)
		if !ok {
			return []common.LinkEntity{}, errors.New("field url does not exist in search result")
		}
		title, ok := e.Properties["title"].(string)
		if !ok {
			return []common.LinkEntity{}, errors.New("field title does not exist in search result")
		}
		description, ok := e.Properties["description"].(string)
		if !ok {
			return []common.LinkEntity{}, errors.New("field description does not exist in search result")
		}
		userIDStr, ok := e.Properties["user_id"].(string)
		if !ok {
			return []common.LinkEntity{}, errors.New("field user_id does not exist in search result")
		}
		resUserID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return []common.LinkEntity{}, err
		}
		if resUserID != userID {
			continue
		}
		folderIDStr, ok := e.Properties["folder_id"].(string)
		if !ok {
			return []common.LinkEntity{}, errors.New("field folder_id does not exist in search result")
		}
		folderID, err := strconv.Atoi(folderIDStr)
		if err != nil {
			return []common.LinkEntity{}, err
		}
		ll = append(ll, common.LinkEntity{
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

func (r *repository) GetOneLinkByID(id int) (common.LinkEntity, error) {
	stmt, err := r.db.Preparex(QueryGetOneLinkByID)
	if err != nil {
		return common.LinkEntity{}, err
	}

	var link common.LinkEntity
	if err := stmt.Get(&link, id); err != nil {
		return common.LinkEntity{}, err
	}

	return link, nil
}

func (r *repository) GetManyLinksInsideDefaultFolder(
	userID int,
	limit int,
	offset int,
	orderBy string,
	sort string,
) ([]common.LinkEntity, error) {
	var q string = QueryGetManyLinksInsideDefaultFolder_OrderByUpdatedAtSortDESC
	switch orderBy {
	case OrderByCreatedAt:
		switch sort {
		case SortASC:
			q = QueryGetManyLinksInsideDefaultFolder_OrderByCreatedAtSortASC
			break
		case SortDESC:
			q = QueryGetManyLinksInsideDefaultFolder_OrderByCreatedAtSortDESC
			break
		}
		break
	case OrderByUpdatedAt:
		switch sort {
		case SortASC:
			q = QueryGetManyLinksInsideDefaultFolder_OrderByUpdatedAtSortASC
			break
		case SortDESC:
			q = QueryGetManyLinksInsideDefaultFolder_OrderByUpdatedAtSortDESC
			break
		}
		break
	}

	stmt, err := r.db.Preparex(q)
	if err != nil {
		return []common.LinkEntity{}, err
	}

	var link []common.LinkEntity
	if err := stmt.Select(&link, userID, limit, offset); err != nil {
		return []common.LinkEntity{}, err
	}

	return link, nil
}

func (r *repository) GetManyLinksInsideFolder(
	userID int,
	folderId int,
	limit int,
	offset int,
	orderBy string,
	sort string,
) ([]common.LinkEntity, error) {
	var q string
	switch orderBy {
	case OrderByCreatedAt:
		switch sort {
		case SortASC:
			q = QueryGetManyLinksInsideFolder_OrderByCreatedAtSortASC
			break
		case SortDESC:
			q = QueryGetManyLinksInsideFolder_OrderByCreatedAtSortDESC
			break
		}
		break
	case OrderByUpdatedAt:
		switch sort {
		case SortASC:
			q = QueryGetManyLinksInsideFolder_OrderByUpdatedAtSortASC
			break
		case SortDESC:
			q = QueryGetManyLinksInsideFolder_OrderByUpdatedAtSortDESC
			break
		}
		break
	}

	stmt, err := r.db.Preparex(q)
	if err != nil {
		return []common.LinkEntity{}, err
	}

	var link []common.LinkEntity
	if err := stmt.Select(&link, userID, folderId, limit, offset); err != nil {
		return []common.LinkEntity{}, err
	}

	return link, nil
}

func (r *repository) UpdateOneLinkByID(
	id int,
	url string,
	title string,
	description string,
	user_id int,
	folder_id int,
) (common.LinkEntity, error) {
	stmt, err := r.db.Preparex(QueryUpdateOneLinkByID)
	if err != nil {
		return common.LinkEntity{}, err
	}

	var link common.LinkEntity
	if err := stmt.Get(&link, id, url, title, description, user_id, folder_id); err != nil {
		return common.LinkEntity{}, err
	}

	return link, nil
}

func (r *repository) DeleteOneLinkByID(id int) (common.LinkEntity, error) {
	stmt, err := r.db.Preparex(QueryDeleteOneLinkByID)
	if err != nil {
		return common.LinkEntity{}, err
	}

	var link common.LinkEntity
	if err := stmt.Get(&link, id); err != nil {
		return common.LinkEntity{}, err
	}

	return link, nil
}
