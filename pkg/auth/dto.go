package auth

type LogInDTO struct {
	Username string `schema:"username,required"`
	Password string `schema:"password,required"`
}
