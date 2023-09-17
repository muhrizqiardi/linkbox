package link

import (
	"github.com/jmoiron/sqlx"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/common"
)

const (
	OrderByCreatedAt = "created_at"
	OrderByUpdatedAt = "updated_at"

	SortASC  = "asc"
	SortDESC = "desc"
)

type Repository interface {
	CreateLink(
		url string,
		title string,
		description string,
		user_id int,
		folder_id int,
	) (common.LinkEntity, error)
	GetOneLinkByID(id int) (common.LinkEntity, error)
	GetManyLinksInsideDefaultFolder(
		userID int,
		limit int,
		offset int,
		orderBy string,
		sort string,
	) ([]common.LinkEntity, error)
	GetManyLinksInsideFolder(
		userID int,
		folder_id int,
		limit int,
		offset int,
		orderBy string,
		sort string,
	) ([]common.LinkEntity, error)
	UpdateOneLinkByID(
		id int,
		url string,
		title string,
		description string,
		user_id int,
		folder_id int,
	) (common.LinkEntity, error)
	DeleteOneLinkByID(id int) (common.LinkEntity, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db}
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
