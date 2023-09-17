package folder

import (
	"errors"
	"regexp"

	"github.com/muhrizqiardi/linkbox/linkbox/pkg/common"
)

var FolderUniqueNameRegex = regexp.MustCompile("^[a-z0-9_]{3,21}$")

var ErrInvalidFolderUniqueName error = errors.New("A folder name can only contains alphanumeric character and underscore, and can only have at least 3 characters and 21 characters maximum")
var ErrInvalidOrderBy error = errors.New("Can only order by `created_at` or `updated_at`")
var ErrInvalidSortMethod error = errors.New("Sort method should be `asc` or `desc`")

type service struct {
	repo common.FolderRepository
}

func NewService(repo common.FolderRepository) *service {
	return &service{repo}
}

func (s *service) Create(payload common.CreateFolderDTO) (common.FolderEntity, error) {
	if !FolderUniqueNameRegex.MatchString(payload.UniqueName) {
		return common.FolderEntity{}, ErrInvalidFolderUniqueName
	}

	folder, err := s.repo.CreateFolder(payload.UniqueName, payload.UserID)
	if err != nil {
		return common.FolderEntity{}, err
	}

	return folder, nil
}

func (s *service) GetOneByID(id int) (common.FolderEntity, error) {
	folder, err := s.repo.GetOneFolderByID(id)
	if err != nil {
		return common.FolderEntity{}, err
	}

	return folder, nil
}

func (s *service) GetOneByUniqueName(uniqueName string, userID int) (common.FolderEntity, error) {
	if !FolderUniqueNameRegex.MatchString(uniqueName) {
		return common.FolderEntity{}, ErrInvalidFolderUniqueName
	}

	folder, err := s.repo.CreateFolder(uniqueName, userID)
	if err != nil {
		return common.FolderEntity{}, err
	}

	return folder, nil
}

func (s *service) GetMany(userID int, options common.GetManyFoldersDTO) ([]common.FolderEntity, error) {
	switch options.OrderBy {
	case GetManyFoldersOrderByCreatedAt:
		switch options.Sort {
		case GetManyFoldersSortAscending:
		case GetManyFoldersSortDescending:
			break
		default:
			return []common.FolderEntity{}, ErrInvalidSortMethod
		}
		break
	case GetManyFoldersOrderByUpdatedAt:
		switch options.Sort {
		case GetManyFoldersSortAscending:
		case GetManyFoldersSortDescending:
			break
		default:
			return []common.FolderEntity{}, ErrInvalidSortMethod
		}
		break
	default:
		return []common.FolderEntity{}, ErrInvalidOrderBy
	}

	folders, err := s.repo.GetManyFolders(options.Limit, options.Offset, options.Sort, options.OrderBy, userID)
	if err != nil {
		return []common.FolderEntity{}, err
	}

	return folders, nil
}

func (s *service) UpdateOneByID(id int, payload common.UpdateFolderDTO) (common.FolderEntity, error) {
	if !FolderUniqueNameRegex.MatchString(payload.UniqueName) {
		return common.FolderEntity{}, ErrInvalidFolderUniqueName
	}

	folder, err := s.repo.UpdateFolderByID(id, payload.UniqueName)
	if err != nil {
		return common.FolderEntity{}, err
	}

	return folder, nil
}

func (s *service) DeleteOneByID(id int) (common.FolderEntity, error) {
	folder, err := s.repo.DeleteFolderByID(id)
	if err != nil {
		return common.FolderEntity{}, err
	}

	return folder, nil
}
