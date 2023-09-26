package service

import (
	"regexp"

	"github.com/muhrizqiardi/linkbox/internal/constant"
	"github.com/muhrizqiardi/linkbox/internal/entities/request"
	"github.com/muhrizqiardi/linkbox/internal/model"
	"github.com/muhrizqiardi/linkbox/internal/repository"
)

const (
	GetManyFoldersOrderByCreatedAt = "created_at"
	GetManyFoldersOrderByUpdatedAt = "updated_at"
	GetManyFoldersSortAscending    = "asc"
	GetManyFoldersSortDescending   = "desc"
)

var FolderUniqueNameRegex = regexp.MustCompile("^[a-z0-9_]{3,21}$")

type FolderService interface {
	Create(payload request.CreateFolderRequest) (model.FolderModel, error)
	GetOneByID(id int) (model.FolderModel, error)
	GetOneByUniqueName(uniqueName string, userID int) (model.FolderModel, error)
	GetMany(userID int, options request.GetManyFoldersRequest) ([]model.FolderModel, error)
	UpdateOneByID(id int, payload request.UpdateFolderRequest) (model.FolderModel, error)
	DeleteOneByID(id int) (model.FolderModel, error)
}

type folderService struct {
	repo repository.FolderRepository
}

func NewFolderService(repo repository.FolderRepository) *folderService {
	return &folderService{repo}
}

func (s *folderService) Create(payload request.CreateFolderRequest) (model.FolderModel, error) {
	if !FolderUniqueNameRegex.MatchString(payload.UniqueName) {
		return model.FolderModel{}, constant.ErrInvalidFolderUniqueName
	}

	folder, err := s.repo.CreateFolder(payload.UniqueName, payload.UserID)
	if err != nil {
		return model.FolderModel{}, err
	}

	return folder, nil
}

func (s *folderService) GetOneByID(id int) (model.FolderModel, error) {
	folder, err := s.repo.GetOneFolderByID(id)
	if err != nil {
		return model.FolderModel{}, err
	}

	return folder, nil
}

func (s *folderService) GetOneByUniqueName(uniqueName string, userID int) (model.FolderModel, error) {
	if !FolderUniqueNameRegex.MatchString(uniqueName) {
		return model.FolderModel{}, constant.ErrInvalidFolderUniqueName
	}

	folder, err := s.repo.CreateFolder(uniqueName, userID)
	if err != nil {
		return model.FolderModel{}, err
	}

	return folder, nil
}

func (s *folderService) GetMany(userID int, options request.GetManyFoldersRequest) ([]model.FolderModel, error) {
	switch options.OrderBy {
	case GetManyFoldersOrderByCreatedAt:
		switch options.Sort {
		case GetManyFoldersSortAscending:
		case GetManyFoldersSortDescending:
			break
		default:
			return []model.FolderModel{}, constant.ErrInvalidGetManyFoldersSortMethod
		}
		break
	case GetManyFoldersOrderByUpdatedAt:
		switch options.Sort {
		case GetManyFoldersSortAscending:
		case GetManyFoldersSortDescending:
			break
		default:
			return []model.FolderModel{}, constant.ErrInvalidGetManyFoldersSortMethod
		}
		break
	default:
		return []model.FolderModel{}, constant.ErrInvalidGetManyFoldersOrderBy
	}

	folders, err := s.repo.GetManyFolders(options.Limit, options.Offset, options.Sort, options.OrderBy, userID)
	if err != nil {
		return []model.FolderModel{}, err
	}

	return folders, nil
}

func (s *folderService) UpdateOneByID(id int, payload request.UpdateFolderRequest) (model.FolderModel, error) {
	if !FolderUniqueNameRegex.MatchString(payload.UniqueName) {
		return model.FolderModel{}, constant.ErrInvalidFolderUniqueName
	}

	folder, err := s.repo.UpdateFolderByID(id, payload.UniqueName)
	if err != nil {
		return model.FolderModel{}, err
	}

	return folder, nil
}

func (s *folderService) DeleteOneByID(id int) (model.FolderModel, error) {
	folder, err := s.repo.DeleteFolderByID(id)
	if err != nil {
		return model.FolderModel{}, err
	}

	return folder, nil
}
