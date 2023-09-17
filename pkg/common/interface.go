package common

import "io"

type Templater interface {
	RegisterPage(w io.Writer, data RegisterPageData) error
	LogInPage(w io.Writer, data LogInPageData) error
	IndexPage(w io.Writer, data IndexPageData) error
	LinksInFolderPage(w io.Writer, data LinksInFolderPageData) error
	LinkFragment(w io.Writer, data LinkFragmentData) error
	EditLinkModalFragment(w io.Writer, data EditLinkModalFragmentData) error
}

type AuthService interface {
	LogIn(payload LogInDTO) (string, error)
	CheckIsValid(token string) (TokenClaims, string, error)
}

type FolderRepository interface {
	CreateFolder(uniqueName string, userID int) (FolderEntity, error)
	GetOneFolderByID(id int) (FolderEntity, error)
	GetOneFolderByUniqueName(uniqueName string, userID int) (FolderEntity, error)
	GetManyFolders(limit int, offset int, sort string, orderBy string, userID int) ([]FolderEntity, error)
	UpdateFolderByID(id int, uniqueName string) (FolderEntity, error)
	DeleteFolderByID(id int) (FolderEntity, error)
}

type FolderService interface {
	Create(payload CreateFolderDTO) (FolderEntity, error)
	GetOneByID(id int) (FolderEntity, error)
	GetOneByUniqueName(uniqueName string, userID int) (FolderEntity, error)
	GetMany(userID int, options GetManyFoldersDTO) ([]FolderEntity, error)
	UpdateOneByID(id int, payload UpdateFolderDTO) (FolderEntity, error)
	DeleteOneByID(id int) (FolderEntity, error)
}

type LinkService interface {
	Create(payload CreateLinkDTO) (LinkEntity, error)
	GetOneByID(id int) (LinkEntity, error)
	GetManyInsideDefaultFolder(userID int, payload GetManyLinksInsideFolderDTO) ([]LinkEntity, error)
	GetManyInsideFolder(userID int, folderId int, payload GetManyLinksInsideFolderDTO) ([]LinkEntity, error)
	UpdateOneByID(id int, payload UpdateLinkDTO) (LinkEntity, error)
	DeleteOneByID(id int) (LinkEntity, error)
}

type UserService interface {
	Create(payload CreateUserDTO) (UserEntity, error)
	GetOneByID(id int) (UserEntity, error)
	GetOneByUsername(username string) (UserEntity, error)
	UpdateOneByID(id int, payload UpdateUserDTO) (UserEntity, error)
	DeleteOneByID(id int) (UserEntity, error)
}
