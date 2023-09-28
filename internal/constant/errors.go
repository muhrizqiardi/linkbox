package constant

import "errors"

var (
	// auth
	ErrLogInUser    = errors.New("Failed to log in to account, make sure your username and password are correct")
	ErrInvalidToken = errors.New("Invalid token")

	// user
	ErrInvalidUsername           = errors.New("Username can only contains alphanumeric character and underscore, and can only have at least 3 characters and 21 characters maximum")
	ErrConfirmPasswordNotMatched = errors.New("Confirm Password field should be equal to Password")
	ErrDuplicateUsername         = errors.New("Username already exists")
	ErrRegisterUser              = errors.New("Failed to register account")
	ErrUserNotFound              = errors.New("User not found")

	// folder
	ErrInvalidFolderUniqueName         = errors.New("A folder name can only contains alphanumeric character and underscore, and can only have at least 3 characters and 21 characters maximum")
	ErrInvalidGetManyFoldersOrderBy    = errors.New("\"orderBy\" should either be \"created_at\" or \"update_at\"")
	ErrInvalidGetManyFoldersSortMethod = errors.New("\"sort\" should either be \"asc\" \"desc\"")

	ErrInvalidGetManyLinksOrderBy    = errors.New("\"orderBy\" should either be \"created_at\" or \"update_at\"")
	ErrInvalidGetManyLinksSortMethod = errors.New("\"sort\" should either be \"asc\" \"desc\"")
)
