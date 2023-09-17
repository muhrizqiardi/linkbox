package link

import (
	"errors"

	"github.com/muhrizqiardi/linkbox/pkg/common"
)

var ErrInvalidOrderBy = errors.New("\"orderBy\" should either be \"created_at\" or \"update_at\"")
var ErrInvalidSortMethod = errors.New("\"sort\" should either be \"asc\" \"desc\"")

type service struct {
	repo common.LinkRepository
}

func NewService(repo common.LinkRepository) *service {
	return &service{repo}
}

func (s *service) Create(payload common.CreateLinkDTO) (common.LinkEntity, error) {
	// TODO: validate payload
	link, err := s.repo.CreateLink(
		payload.URL,
		payload.Title,
		payload.Description,
		payload.UserID,
		payload.FolderID,
	)
	if err != nil {
		return common.LinkEntity{}, err
	}

	return link, nil
}

func (s *service) GetOneByID(id int) (common.LinkEntity, error) {
	link, err := s.repo.GetOneLinkByID(id)
	if err != nil {
		return common.LinkEntity{}, err
	}

	return link, nil
}

func (s *service) GetManyInsideDefaultFolder(userID int, payload common.GetManyLinksInsideFolderDTO) ([]common.LinkEntity, error) {
	switch payload.OrderBy {
	case OrderByCreatedAt:
		switch payload.Sort {
		case SortASC:
		case SortDESC:
			break
		default:
			return []common.LinkEntity{}, ErrInvalidSortMethod
		}
	case OrderByUpdatedAt:
		switch payload.Sort {
		case SortASC:
		case SortDESC:
			break
		default:
			return []common.LinkEntity{}, ErrInvalidSortMethod
		}
	default:
		return []common.LinkEntity{}, ErrInvalidOrderBy
	}

	links, err := s.repo.GetManyLinksInsideDefaultFolder(
		userID,
		payload.Limit,
		payload.Offset,
		payload.OrderBy,
		payload.Sort,
	)
	if err != nil {
		return []common.LinkEntity{}, err
	}

	return links, nil
}

func (s *service) GetManyInsideFolder(userID int, folderId int, payload common.GetManyLinksInsideFolderDTO) ([]common.LinkEntity, error) {
	switch payload.OrderBy {
	case OrderByCreatedAt:
		switch payload.Sort {
		case SortASC:
		case SortDESC:
			break
		default:
			return []common.LinkEntity{}, ErrInvalidSortMethod
		}
	case OrderByUpdatedAt:
		switch payload.Sort {
		case SortASC:
		case SortDESC:
			break
		default:
			return []common.LinkEntity{}, ErrInvalidSortMethod
		}
	default:
		return []common.LinkEntity{}, ErrInvalidOrderBy
	}

	links, err := s.repo.GetManyLinksInsideFolder(
		userID,
		folderId,
		payload.Limit,
		payload.Offset,
		payload.OrderBy,
		payload.Sort,
	)
	if err != nil {
		return []common.LinkEntity{}, err
	}

	return links, nil
}

func (s *service) UpdateOneByID(id int, payload common.UpdateLinkDTO) (common.LinkEntity, error) {
	// TODO: validate payload
	link, err := s.repo.UpdateOneLinkByID(
		id,
		payload.URL,
		payload.Title,
		payload.Description,
		payload.UserID,
		payload.FolderID,
	)
	if err != nil {
		return common.LinkEntity{}, err
	}

	return link, nil

}

func (s *service) DeleteOneByID(id int) (common.LinkEntity, error) {
	link, err := s.repo.DeleteOneLinkByID(id)
	if err != nil {
		return common.LinkEntity{}, err
	}

	return link, nil
}
