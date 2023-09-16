package link

import (
	"errors"
)

var ErrInvalidOrderBy = errors.New("\"orderBy\" should either be \"created_at\" or \"update_at\"")
var ErrInvalidSortMethod = errors.New("\"sort\" should either be \"asc\" \"desc\"")

type Service interface {
	Create(payload CreateLinkDTO) (LinkEntity, error)
	GetOneByID(id int) (LinkEntity, error)
	GetManyInsideDefaultFolder(userID int, payload GetManyLinksInsideFolderDTO) ([]LinkEntity, error)
	GetManyInsideFolder(userID int, folderId int, payload GetManyLinksInsideFolderDTO) ([]LinkEntity, error)
	UpdateOneByID(id int, payload UpdateLinkDTO) (LinkEntity, error)
	DeleteOneByID(id int) (LinkEntity, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) Create(payload CreateLinkDTO) (LinkEntity, error) {
	// TODO: validate payload
	link, err := s.repo.CreateLink(
		payload.URL,
		payload.Title,
		payload.Description,
		payload.UserID,
		payload.FolderID,
	)
	if err != nil {
		return LinkEntity{}, err
	}

	return link, nil
}

func (s *service) GetOneByID(id int) (LinkEntity, error) {
	link, err := s.repo.GetOneLinkByID(id)
	if err != nil {
		return LinkEntity{}, err
	}

	return link, nil
}

func (s *service) GetManyInsideDefaultFolder(userID int, payload GetManyLinksInsideFolderDTO) ([]LinkEntity, error) {
	switch payload.OrderBy {
	case OrderByCreatedAt:
		switch payload.Sort {
		case SortASC:
		case SortDESC:
			break
		default:
			return []LinkEntity{}, ErrInvalidSortMethod
		}
	case OrderByUpdatedAt:
		switch payload.Sort {
		case SortASC:
		case SortDESC:
			break
		default:
			return []LinkEntity{}, ErrInvalidSortMethod
		}
	default:
		return []LinkEntity{}, ErrInvalidOrderBy
	}

	links, err := s.repo.GetManyLinksInsideDefaultFolder(
		userID,
		payload.Limit,
		payload.Offset,
		payload.OrderBy,
		payload.Sort,
	)
	if err != nil {
		return []LinkEntity{}, err
	}

	return links, nil
}

func (s *service) GetManyInsideFolder(userID int, folderId int, payload GetManyLinksInsideFolderDTO) ([]LinkEntity, error) {
	switch payload.OrderBy {
	case OrderByCreatedAt:
		switch payload.Sort {
		case SortASC:
		case SortDESC:
			break
		default:
			return []LinkEntity{}, ErrInvalidSortMethod
		}
	case OrderByUpdatedAt:
		switch payload.Sort {
		case SortASC:
		case SortDESC:
			break
		default:
			return []LinkEntity{}, ErrInvalidSortMethod
		}
	default:
		return []LinkEntity{}, ErrInvalidOrderBy
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
		return []LinkEntity{}, err
	}

	return links, nil
}

func (s *service) UpdateOneByID(id int, payload UpdateLinkDTO) (LinkEntity, error) {
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
		return LinkEntity{}, err
	}

	return link, nil

}

func (s *service) DeleteOneByID(id int) (LinkEntity, error) {
	link, err := s.repo.DeleteOneLinkByID(id)
	if err != nil {
		return LinkEntity{}, err
	}

	return link, nil
}
