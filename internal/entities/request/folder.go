package request

type CreateFolderRequest struct {
	UniqueName string `schema:"uniqueName"`
	UserID     int    `schema:"userId"`
}

type GetManyFoldersRequest struct {
	OrderBy string `schema:"orderBy"`
	Sort    string `schema:"sort"`
	Limit   int    `schema:"limit"`
	Offset  int    `schema:"offset"`
}

type UpdateFolderRequest struct {
	UniqueName string `schema:"uniqueName"`
}
