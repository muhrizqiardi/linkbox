package link

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
