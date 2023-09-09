package folder

import (
	"errors"
	"regexp"
)

var FolderUniqueNameRegex = regexp.MustCompile("^[a-z0-9_]{3,21}$")

var ErrInvalidFolderUniqueName error = errors.New("A folder name can only contains alphanumeric character and underscore, and can only have at least 3 characters and 21 characters maximum")
var ErrInvalidOrderBy error = errors.New("Can only order by `created_at` or `updated_at`")
var ErrInvalidSortMethod error = errors.New("Sort method should be `asc` or `desc`")

type Service interface {
	Create(payload CreateFolderDTO) (FolderEntity, error)
	GetOneByID(id int) (FolderEntity, error)
	GetOneByUniqueName(uniqueName string, userID int) (FolderEntity, error)
	GetMany(userID int, options GetManyFoldersDTO) ([]FolderEntity, error)
	UpdateOneByID(id int, payload UpdateFolderDTO) (FolderEntity, error)
	DeleteOneByID(id int) (FolderEntity, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) Create(payload CreateFolderDTO) (FolderEntity, error) {
	if !FolderUniqueNameRegex.MatchString(payload.UniqueName) {
		return FolderEntity{}, ErrInvalidFolderUniqueName
	}

	folder, err := s.repo.CreateFolder(payload.UniqueName, payload.UserID)
	if err != nil {
		return FolderEntity{}, err
	}

	return folder, nil
}

func (s *service) GetOneByID(id int) (FolderEntity, error) {
	folder, err := s.repo.GetOneFolderByID(id)
	if err != nil {
		return FolderEntity{}, err
	}

	return folder, nil
}

func (s *service) GetOneByUniqueName(uniqueName string, userID int) (FolderEntity, error) {
	if !FolderUniqueNameRegex.MatchString(uniqueName) {
		return FolderEntity{}, ErrInvalidFolderUniqueName
	}

	folder, err := s.repo.CreateFolder(uniqueName, userID)
	if err != nil {
		return FolderEntity{}, err
	}

	return folder, nil
}

func (s *service) GetMany(userID int, options GetManyFoldersDTO) ([]FolderEntity, error) {
	switch options.OrderBy {
	case GetManyFoldersOrderByCreatedAt:
		switch options.Sort {
		case GetManyFoldersSortAscending:
		case GetManyFoldersSortDescending:
			break
		default:
			return []FolderEntity{}, ErrInvalidSortMethod
		}
		break
	case GetManyFoldersOrderByUpdatedAt:
		switch options.Sort {
		case GetManyFoldersSortAscending:
		case GetManyFoldersSortDescending:
			break
		default:
			return []FolderEntity{}, ErrInvalidSortMethod
		}
		break
	default:
		return []FolderEntity{}, ErrInvalidOrderBy
	}

	folders, err := s.repo.GetManyFolders(options.Limit, options.Offset, options.Sort, options.OrderBy, userID)
	if err != nil {
		return []FolderEntity{}, err
	}

	return folders, nil
}

func (s *service) UpdateOneByID(id int, payload UpdateFolderDTO) (FolderEntity, error) {
	if !FolderUniqueNameRegex.MatchString(payload.UniqueName) {
		return FolderEntity{}, ErrInvalidFolderUniqueName
	}

	folder, err := s.repo.UpdateFolderByID(id, payload.UniqueName)
	if err != nil {
		return FolderEntity{}, err
	}

	return folder, nil
}

func (s *service) DeleteOneByID(id int) (FolderEntity, error) {
	folder, err := s.repo.DeleteFolderByID(id)
	if err != nil {
		return FolderEntity{}, err
	}

	return folder, nil
}
