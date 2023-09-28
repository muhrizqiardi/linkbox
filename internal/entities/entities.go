package entities

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/muhrizqiardi/linkbox/internal/entities/request"
	"github.com/muhrizqiardi/linkbox/internal/model"
)

type TokenClaims struct {
	UserID int `json:"userId"`
	jwt.RegisteredClaims
}

type PageMetaData struct {
	Title       string
	Description string
	ImageURL    string
}

type RegisterPageData struct {
	Errors []string
	PageMetaData
}

type LogInPageData struct {
	Errors []string
	PageMetaData
}

type IndexPageData struct {
	User    model.UserModel
	Folders []model.FolderModel
	Links   []model.LinkModel
	PageMetaData
}

type SearchPageData struct {
	User    model.UserModel
	Folders []model.FolderModel
	Links   []model.LinkModel
	PageMetaData
}

type SearchResultsFragmentData struct {
	Links []model.LinkModel
	PageMetaData
}

type LinksInFolderPageData struct {
	User    model.UserModel
	Folder  model.FolderModel
	Folders []model.FolderModel
	Links   []model.LinkModel
	PageMetaData
}

type LinkFragmentData struct {
	Link model.LinkModel
}

type NewLinkModalFragmentData struct {
	User             model.UserModel
	Folders          []model.FolderModel
	InitialFormValue request.CreateLinkRequest
	Errors           []string
}

type EditLinkModalFragmentData struct {
	User    model.UserModel
	Folders []model.FolderModel
	Link    model.LinkModel
	Errors  []string
}

type DeleteLinkConfirmationModalFragmentData struct {
	LinkID int
}
