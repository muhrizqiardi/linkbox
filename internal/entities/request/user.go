package request

type CreateUserRequest struct {
	Username        string `schema:"username,required"`
	Password        string `schema:"password,required"`
	ConfirmPassword string `schema:"confirmPassword,required"`
}

type UpdateUserRequest struct {
	Username        string `schema:"username,required"`
	Password        string `schema:"password,required"`
	ConfirmPassword string `schema:"confirmPassword,required"`
}
