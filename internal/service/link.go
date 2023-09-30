package service

import (
	"fmt"
	"net/http"

	"github.com/muhrizqiardi/linkbox/internal/constant"
	"github.com/muhrizqiardi/linkbox/internal/entities"
	"github.com/muhrizqiardi/linkbox/internal/entities/request"
	"github.com/muhrizqiardi/linkbox/internal/entities/response"
	"github.com/muhrizqiardi/linkbox/internal/model"
	"github.com/muhrizqiardi/linkbox/internal/repository"
	"github.com/muhrizqiardi/linkbox/internal/util"
)

type LinkService interface {
	Create(payload request.CreateLinkRequest) (model.LinkModel, error)
	GetOneByID(id int) (model.LinkModel, error)
	SearchFullText(userID int, query string) ([]model.LinkModel, error)
	GetManyInsideDefaultFolder(userID int, payload request.GetManyLinksInsideFolderRequest) ([]response.LinkWithMediaResponse, error)
	GetManyInsideFolder(userID int, folderId int, payload request.GetManyLinksInsideFolderRequest) ([]response.LinkWithMediaResponse, error)
	UpdateOneByID(id int, payload request.UpdateLinkRequest) (model.LinkModel, error)
	DeleteOneByID(id int) (model.LinkModel, error)
}

type linkService struct {
	lr  repository.LinkRepository
	lmr repository.LinkMediaRepository
}

func NewLinkService(lr repository.LinkRepository, lmr repository.LinkMediaRepository) *linkService {
	return &linkService{lr, lmr}
}

func (ls *linkService) Create(payload request.CreateLinkRequest) (model.LinkModel, error) {
	// TODO: validate payload
	metadata := entities.LinkMetadata{
		OG: entities.OpenGraph{
			Title:       payload.Title,
			URL:         payload.URL,
			Description: payload.Description,
			OGImage:     []entities.OGImage{},
		},
	}
	if payload.Title == "" || payload.Description == "" {
		req, _ := http.NewRequest(http.MethodGet, payload.URL, nil)
		res, err := http.DefaultClient.Do(req)
		if err == nil {
			metadata, err = util.MetadataScraper(res)
			if err != nil {
				fmt.Println("failed to scrape metadata:", err)
			}
		}
	}

	link, err := ls.lr.CreateLink(
		payload.URL,
		metadata.OG.Title,
		metadata.OG.Description,
		payload.UserID,
		payload.FolderID,
	)
	if err != nil {
		return model.LinkModel{}, err
	}

	if len(metadata.OG.OGImage) > 0 {
		for _, e := range metadata.OG.OGImage {
			if _, err := ls.lmr.Insert(link.ID, e.URL); err != nil {
				fmt.Println("failed to insert link media:", err)
			}
		}
	}

	return link, nil
}

func (ls *linkService) GetOneByID(id int) (model.LinkModel, error) {
	link, err := ls.lr.GetOneLinkByID(id)
	if err != nil {
		return model.LinkModel{}, err
	}

	return link, nil
}

func (s *linkService) SearchFullText(userID int, query string) ([]model.LinkModel, error) {
	ll, err := s.lr.SearchFullText(userID, query)
	if err != nil {
		return []model.LinkModel{}, err
	}

	return ll, nil
}

func (ls *linkService) GetManyInsideDefaultFolder(userID int, payload request.GetManyLinksInsideFolderRequest) ([]response.LinkWithMediaResponse, error) {
	switch payload.OrderBy {
	case constant.GetManyLinksOrderByCreatedAt:
		switch payload.Sort {
		case constant.GetManyLinksSortASC:
		case constant.GetManyLinksSortDESC:
			break
		default:
			return []response.LinkWithMediaResponse{}, constant.ErrInvalidGetManyLinksSortMethod
		}
	case constant.GetManyLinksOrderByUpdatedAt:
		switch payload.Sort {
		case constant.GetManyLinksSortASC:
		case constant.GetManyLinksSortDESC:
			break
		default:
			return []response.LinkWithMediaResponse{}, constant.ErrInvalidGetManyLinksSortMethod
		}
	default:
		return []response.LinkWithMediaResponse{}, constant.ErrInvalidGetManyFoldersOrderBy
	}

	links, err := ls.lr.GetManyLinksInsideDefaultFolder(
		userID,
		payload.Limit,
		payload.Offset,
		payload.OrderBy,
		payload.Sort,
	)
	if err != nil {
		return []response.LinkWithMediaResponse{}, err
	}

	return links, nil
}

func (ls *linkService) GetManyInsideFolder(userID int, folderId int, payload request.GetManyLinksInsideFolderRequest) ([]response.LinkWithMediaResponse, error) {
	switch payload.OrderBy {
	case constant.GetManyLinksOrderByCreatedAt:
		switch payload.Sort {
		case constant.GetManyLinksSortASC:
		case constant.GetManyLinksSortDESC:
			break
		default:
			return []response.LinkWithMediaResponse{}, constant.ErrInvalidGetManyLinksSortMethod
		}
	case constant.GetManyLinksOrderByUpdatedAt:
		switch payload.Sort {
		case constant.GetManyLinksSortASC:
		case constant.GetManyLinksSortDESC:
			break
		default:
			return []response.LinkWithMediaResponse{}, constant.ErrInvalidGetManyLinksSortMethod
		}
	default:
		return []response.LinkWithMediaResponse{}, constant.ErrInvalidGetManyFoldersOrderBy
	}

	links, err := ls.lr.GetManyLinksInsideFolder(
		userID,
		folderId,
		payload.Limit,
		payload.Offset,
		payload.OrderBy,
		payload.Sort,
	)
	if err != nil {
		return []response.LinkWithMediaResponse{}, err
	}

	return links, nil
}

func (ls *linkService) UpdateOneByID(id int, payload request.UpdateLinkRequest) (model.LinkModel, error) {
	// TODO: validate payload
	link, err := ls.lr.UpdateOneLinkByID(
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
	link, err := ls.lr.DeleteOneLinkByID(id)
	if err != nil {
		return model.LinkModel{}, err
	}

	return link, nil
}
