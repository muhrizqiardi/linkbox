package common

type LogInDTO struct {
	Username string `schema:"username,required"`
	Password string `schema:"password,required"`
}

type CreateFolderDTO struct {
	UniqueName string `schema:"uniqueName"`
	UserID     int    `schema:"userId"`
}

type GetManyFoldersDTO struct {
	OrderBy string `schema:"orderBy"`
	Sort    string `schema:"sort"`
	Limit   int    `schema:"limit"`
	Offset  int    `schema:"offset"`
}

type UpdateFolderDTO struct {
	UniqueName string `schema:"uniqueName"`
}

type CreateLinkDTO struct {
	URL         string `schema:"url"`
	Title       string `schema:"title"`
	Description string `schema:"description"`
	UserID      int    `schema:"userId"`
	FolderID    int    `schema:"folderId"`
}

type GetManyLinksInsideFolderDTO struct {
	Limit   int
	Offset  int
	OrderBy string
	Sort    string
}

type UpdateLinkDTO struct {
	URL         string
	Title       string
	Description string
	UserID      int
	FolderID    int
}

type CreateUserDTO struct {
	Username        string `schema:"username,required"`
	Password        string `schema:"password,required"`
	ConfirmPassword string `schema:"confirmPassword,required"`
}

type UpdateUserDTO struct {
	Username        string `schema:"username,required"`
	Password        string `schema:"password,required"`
	ConfirmPassword string `schema:"confirmPassword,required"`
}
