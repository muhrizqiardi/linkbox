package entities

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/muhrizqiardi/linkbox/internal/entities/request"
	"github.com/muhrizqiardi/linkbox/internal/entities/response"
	"github.com/muhrizqiardi/linkbox/internal/model"
)

type OGImage struct {
	URL       string
	SecureURL string
	Type      string
	Width     string
	Height    string
	Alt       string
}

type OGVideo struct {
	URL       string
	SecureURL string
	Type      string
	Width     string
	Height    string
	Alt       string
}

type OGAudio struct {
	URL       string
	SecureURL string
	Type      string
}

type OpenGraph struct {
	Title       string
	Type        string
	URL         string
	Description string
	OGImage     []OGImage
	OGVideo     []OGVideo
	OGAudio     []OGAudio
}

type LinkMetadata struct {
	OG OpenGraph
}

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
	User     model.UserModel
	Folders  []model.FolderModel
	Links    []response.LinkWithMediaResponse
	FolderID int
	NextPage int
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

type NewFolderModalFragmentData struct {
	User model.UserModel
}

type LinksInFolderPageData struct {
	User     model.UserModel
	Folder   model.FolderModel
	Folders  []model.FolderModel
	Links    []response.LinkWithMediaResponse
	FolderID int
	NextPage int
	PageMetaData
}

type LinksFragmentData struct {
	Links    []response.LinkWithMediaResponse
	FolderID int
	NextPage int
}

type LinkFragmentData struct {
	Link response.LinkWithMediaResponse
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
