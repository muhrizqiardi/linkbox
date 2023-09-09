package user

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
