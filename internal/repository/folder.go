package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/muhrizqiardi/linkbox/internal/model"
	"github.com/muhrizqiardi/linkbox/internal/query"
	// "github.com/muhrizqiardi/linkbox/pkg/common"
)

const (
	GetManyFoldersOrderByCreatedAt = "created_at"
	GetManyFoldersOrderByUpdatedAt = "updated_at"
	GetManyFoldersSortAscending    = "asc"
	GetManyFoldersSortDescending   = "desc"
)

type FolderRepository interface {
	CreateFolder(uniqueName string, userID int) (model.FolderModel, error)
	GetOneFolderByID(id int) (model.FolderModel, error)
	GetOneFolderByUniqueName(uniqueName string, userID int) (model.FolderModel, error)
	GetManyFolders(limit int, offset int, sort string, orderBy string, userID int) ([]model.FolderModel, error)
	UpdateFolderByID(id int, uniqueName string) (model.FolderModel, error)
	DeleteFolderByID(id int) (model.FolderModel, error)
}

type folderRepository struct {
	db *sqlx.DB
}

func NewFolderRepository(db *sqlx.DB) *folderRepository {
	return &folderRepository{db}
}

func (r *folderRepository) CreateFolder(uniqueName string, userID int) (model.FolderModel, error) {
	stmt, err := r.db.Preparex(query.QueryCreateFolder)
	if err != nil {
		return model.FolderModel{}, err
	}

	var folder model.FolderModel
	if err := stmt.Get(&folder, uniqueName, userID); err != nil {
		return model.FolderModel{}, err
	}

	return folder, nil
}

func (r *folderRepository) GetOneFolderByID(id int) (model.FolderModel, error) {
	stmt, err := r.db.Preparex(query.QueryGetOneFolderByID)
	if err != nil {
		return model.FolderModel{}, err
	}

	var folder model.FolderModel
	if err := stmt.Get(&folder, id); err != nil {
		return model.FolderModel{}, err
	}

	return folder, nil
}

func (r *folderRepository) GetOneFolderByUniqueName(uniqueName string, userID int) (model.FolderModel, error) {
	stmt, err := r.db.Preparex(query.QueryGetOneFolderByUniqueName)
	if err != nil {
		return model.FolderModel{}, err
	}

	var folder model.FolderModel
	if err := stmt.Get(&folder, uniqueName, userID); err != nil {
		return model.FolderModel{}, err
	}

	return folder, nil
}

func (r *folderRepository) GetManyFolders(limit int, offset int, sort string, orderBy string, userID int) ([]model.FolderModel, error) {
	var q string
	switch orderBy {
	case GetManyFoldersOrderByCreatedAt:
		switch sort {
		case GetManyFoldersSortAscending:
			q = query.QueryGetManyFolders_ByCreatedAtASC
			break
		case GetManyFoldersSortDescending:
			q = query.QueryGetManyFolders_ByCreatedAtDESC
			break
		default:
			q = query.QueryGetManyFolders_ByCreatedAtDESC
			break
		}
		break
	case GetManyFoldersOrderByUpdatedAt:
		switch sort {
		case GetManyFoldersSortAscending:
			q = query.QueryGetManyFolders_ByUpdatedAtASC
			break
		case GetManyFoldersSortDescending:
			q = query.QueryGetManyFolders_ByUpdatedAtDESC
			break
		default:
			q = query.QueryGetManyFolders_ByUpdatedAtDESC
			break
		}
	default:
		q = query.QueryGetManyFolders_ByUpdatedAtDESC
		break
	}

	stmt, err := r.db.Preparex(q)
	if err != nil {
		return []model.FolderModel{}, err
	}

	var folders []model.FolderModel
	if err := stmt.Select(&folders, limit, offset, userID); err != nil {
		return []model.FolderModel{}, err
	}

	return folders, nil
}

func (r *folderRepository) UpdateFolderByID(id int, uniqueName string) (model.FolderModel, error) {
	stmt, err := r.db.Preparex(query.QueryUpdateOneFolderByID)
	if err != nil {
		return model.FolderModel{}, err
	}

	var folder model.FolderModel
	if err := stmt.Get(&folder, id, uniqueName); err != nil {
		return model.FolderModel{}, err
	}

	return folder, nil
}

func (r *folderRepository) DeleteFolderByID(id int) (model.FolderModel, error) {
	stmt, err := r.db.Preparex(query.QueryDeleteOneFolderByID)
	if err != nil {
		return model.FolderModel{}, err
	}

	var folder model.FolderModel
	if err := stmt.Get(&folder, id); err != nil {
		return model.FolderModel{}, err
	}

	return folder, nil
}
