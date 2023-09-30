package service

import (
	"github.com/muhrizqiardi/linkbox/internal/entities/request"
	"github.com/muhrizqiardi/linkbox/internal/model"
	"github.com/muhrizqiardi/linkbox/internal/repository"
)

type LinkMediaService interface {
	Add(payload request.CreateLinkMediaRequest) (model.LinkMediaModel, error)
}

type linkMediaService struct {
	lmr repository.LinkMediaRepository
}

func NewLinkMediaService(lmr repository.LinkMediaRepository) *linkMediaService {
	return &linkMediaService{lmr}
}
