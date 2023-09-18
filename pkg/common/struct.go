package common

type MetaData struct {
	Title       string
	Description string
	ImageURL    string
}

type RegisterPageData struct {
	Errors []string
	MetaData
}

type LogInPageData struct {
	Errors []string
	MetaData
}

type IndexPageData struct {
	User    UserEntity
	Folders []FolderEntity
	Links   []LinkEntity
	MetaData
}

type SearchPageData struct {
	User    UserEntity
	Folders []FolderEntity
	Links   []LinkEntity
	MetaData
}

type SearchResultsFragmentData struct {
	Links []LinkEntity
	MetaData
}

type LinksInFolderPageData struct {
	User    UserEntity
	Folder  FolderEntity
	Folders []FolderEntity
	Links   []LinkEntity
	MetaData
}

type LinkFragmentData struct {
	Link LinkEntity
}

type EditLinkModalFragmentData struct {
	User    UserEntity
	Folders []FolderEntity
	Link    LinkEntity
}

type DeleteLinkConfirmationModalFragmentData struct {
	LinkID int
}
