package service

import (
	"github.com/muhrizqiardi/linkbox/internal/constant"
	"github.com/muhrizqiardi/linkbox/internal/entities/request"
	"github.com/muhrizqiardi/linkbox/internal/model"
	"github.com/muhrizqiardi/linkbox/internal/repository"
)

type LinkService interface {
	Create(payload request.CreateLinkRequest) (model.LinkModel, error)
	GetOneByID(id int) (model.LinkModel, error)
	SearchFullText(userID int, query string) ([]model.LinkModel, error)
	GetManyInsideDefaultFolder(userID int, payload request.GetManyLinksInsideFolderRequest) ([]model.LinkModel, error)
	GetManyInsideFolder(userID int, folderId int, payload request.GetManyLinksInsideFolderRequest) ([]model.LinkModel, error)
	UpdateOneByID(id int, payload request.UpdateLinkRequest) (model.LinkModel, error)
	DeleteOneByID(id int) (model.LinkModel, error)
}

type linkService struct {
	repo repository.LinkRepository
}

func NewLinkService(repo repository.LinkRepository) *linkService {
	return &linkService{repo}
}

func (ls *linkService) Create(payload request.CreateLinkRequest) (model.LinkModel, error) {
	// TODO: validate payload
	link, err := ls.repo.CreateLink(
		payload.URL,
		payload.Title,
		payload.Description,
		payload.UserID,
		payload.FolderID,
	)
	if err != nil {
		return model.LinkModel{}, err
	}

	return link, nil
}

func (ls *linkService) GetOneByID(id int) (model.LinkModel, error) {
	link, err := ls.repo.GetOneLinkByID(id)
	if err != nil {
		return model.LinkModel{}, err
	}

	return link, nil
}

func (s *linkService) SearchFullText(userID int, query string) ([]model.LinkModel, error) {
	ll, err := s.repo.SearchFullText(userID, query)
	if err != nil {
		return []model.LinkModel{}, err
	}

	return ll, nil
}

func (ls *linkService) GetManyInsideDefaultFolder(userID int, payload request.GetManyLinksInsideFolderRequest) ([]model.LinkModel, error) {
	switch payload.OrderBy {
	case constant.GetManyLinksOrderByCreatedAt:
		switch payload.Sort {
		case constant.GetManyLinksSortASC:
		case constant.GetManyLinksSortDESC:
			break
		default:
			return []model.LinkModel{}, constant.ErrInvalidGetManyLinksSortMethod
		}
	case constant.GetManyLinksOrderByUpdatedAt:
		switch payload.Sort {
		case constant.GetManyLinksSortASC:
		case constant.GetManyLinksSortDESC:
			break
		default:
			return []model.LinkModel{}, constant.ErrInvalidGetManyLinksSortMethod
		}
	default:
		return []model.LinkModel{}, constant.ErrInvalidGetManyFoldersOrderBy
	}

	links, err := ls.repo.GetManyLinksInsideDefaultFolder(
		userID,
		payload.Limit,
		payload.Offset,
		payload.OrderBy,
		payload.Sort,
	)
	if err != nil {
		return []model.LinkModel{}, err
	}

	return links, nil
}

func (ls *linkService) GetManyInsideFolder(userID int, folderId int, payload request.GetManyLinksInsideFolderRequest) ([]model.LinkModel, error) {
	switch payload.OrderBy {
	case constant.GetManyLinksOrderByCreatedAt:
		switch payload.Sort {
		case constant.GetManyLinksSortASC:
		case constant.GetManyLinksSortDESC:
			break
		default:
			return []model.LinkModel{}, constant.ErrInvalidGetManyLinksSortMethod
		}
	case constant.GetManyLinksOrderByUpdatedAt:
		switch payload.Sort {
		case constant.GetManyLinksSortASC:
		case constant.GetManyLinksSortDESC:
			break
		default:
			return []model.LinkModel{}, constant.ErrInvalidGetManyLinksSortMethod
		}
	default:
		return []model.LinkModel{}, constant.ErrInvalidGetManyFoldersOrderBy
	}

	links, err := ls.repo.GetManyLinksInsideFolder(
		userID,
		folderId,
		payload.Limit,
		payload.Offset,
		payload.OrderBy,
		payload.Sort,
	)
	if err != nil {
		return []model.LinkModel{}, err
	}

	return links, nil
}

func (ls *linkService) UpdateOneByID(id int, payload request.UpdateLinkRequest) (model.LinkModel, error) {
	// TODO: validate payload
	link, err := ls.repo.UpdateOneLinkByID(
		id,
		payload.URL,
		payload.Title,
		payload.Description,
		payload.UserID,
		payload.FolderID,
	)
	if err != nil {
		return model.LinkModel{}, err
	}

	return link, nil

}

func (ls *linkService) DeleteOneByID(id int) (model.LinkModel, error) {
	link, err := ls.repo.DeleteOneLinkByID(id)
	if err != nil {
		return model.LinkModel{}, err
	}

	return link, nil
}
