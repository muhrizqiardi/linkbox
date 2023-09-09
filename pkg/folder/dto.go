package folder

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
