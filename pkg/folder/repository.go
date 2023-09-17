package folder

import (
	"github.com/jmoiron/sqlx"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/common"
)

const (
	GetManyFoldersOrderByCreatedAt = "created_at"
	GetManyFoldersOrderByUpdatedAt = "updated_at"
	GetManyFoldersSortAscending    = "asc"
	GetManyFoldersSortDescending   = "desc"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repository {
	return &repository{db}
}

func (r *repository) CreateFolder(uniqueName string, userID int) (common.FolderEntity, error) {
	stmt, err := r.db.Preparex(QueryCreateFolder)
	if err != nil {
		return common.FolderEntity{}, err
	}

	var folder common.FolderEntity
	if err := stmt.Get(&folder, uniqueName, userID); err != nil {
		return common.FolderEntity{}, err
	}

	return folder, nil
}

func (r *repository) GetOneFolderByID(id int) (common.FolderEntity, error) {
	stmt, err := r.db.Preparex(QueryGetOneFolderByID)
	if err != nil {
		return common.FolderEntity{}, err
	}

	var folder common.FolderEntity
	if err := stmt.Get(&folder, id); err != nil {
		return common.FolderEntity{}, err
	}

	return folder, nil
}

func (r *repository) GetOneFolderByUniqueName(uniqueName string, userID int) (common.FolderEntity, error) {
	stmt, err := r.db.Preparex(QueryGetOneFolderByUniqueName)
	if err != nil {
		return common.FolderEntity{}, err
	}

	var folder common.FolderEntity
	if err := stmt.Get(&folder, uniqueName, userID); err != nil {
		return common.FolderEntity{}, err
	}

	return folder, nil
}

func (r *repository) GetManyFolders(limit int, offset int, sort string, orderBy string, userID int) ([]common.FolderEntity, error) {
	var q string
	switch orderBy {
	case GetManyFoldersOrderByCreatedAt:
		switch sort {
		case GetManyFoldersSortAscending:
			q = QueryGetManyFolders_ByCreatedAtASC
			break
		case GetManyFoldersSortDescending:
			q = QueryGetManyFolders_ByCreatedAtDESC
			break
		default:
			q = QueryGetManyFolders_ByCreatedAtDESC
			break
		}
		break
	case GetManyFoldersOrderByUpdatedAt:
		switch sort {
		case GetManyFoldersSortAscending:
			q = QueryGetManyFolders_ByUpdatedAtASC
			break
		case GetManyFoldersSortDescending:
			q = QueryGetManyFolders_ByUpdatedAtDESC
			break
		default:
			q = QueryGetManyFolders_ByUpdatedAtDESC
			break
		}
	default:
		q = QueryGetManyFolders_ByUpdatedAtDESC
		break
	}

	stmt, err := r.db.Preparex(q)
	if err != nil {
		return []common.FolderEntity{}, err
	}

	var folders []common.FolderEntity
	if err := stmt.Select(&folders, limit, offset, userID); err != nil {
		return []common.FolderEntity{}, err
	}

	return folders, nil
}

func (r *repository) UpdateFolderByID(id int, uniqueName string) (common.FolderEntity, error) {
	stmt, err := r.db.Preparex(QueryUpdateOneFolderByID)
	if err != nil {
		return common.FolderEntity{}, err
	}

	var folder common.FolderEntity
	if err := stmt.Get(&folder, id, uniqueName); err != nil {
		return common.FolderEntity{}, err
	}

	return folder, nil
}

func (r *repository) DeleteFolderByID(id int) (common.FolderEntity, error) {
	stmt, err := r.db.Preparex(QueryDeleteOneFolderByID)
	if err != nil {
		return common.FolderEntity{}, err
	}

	var folder common.FolderEntity
	if err := stmt.Get(&folder, id); err != nil {
		return common.FolderEntity{}, err
	}

	return folder, nil
}
