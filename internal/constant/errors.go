package constant

import "errors"

var (
	ErrInvalidUsername         = errors.New("Username can only contains alphanumeric character and underscore, and can only have at least 3 characters and 21 characters maximum")
	ErrInvalidFolderUniqueName = errors.New("A folder name can only contains alphanumeric character and underscore, and can only have at least 3 characters and 21 characters maximum")

	ErrInvalidGetManyFoldersOrderBy    = errors.New("\"orderBy\" should either be \"created_at\" or \"update_at\"")
	ErrInvalidGetManyFoldersSortMethod = errors.New("\"sort\" should either be \"asc\" \"desc\"")

	ErrInvalidGetManyLinksOrderBy    = errors.New("\"orderBy\" should either be \"created_at\" or \"update_at\"")
	ErrInvalidGetManyLinksSortMethod = errors.New("\"sort\" should either be \"asc\" \"desc\"")
)
