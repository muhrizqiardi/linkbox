package request

type LogInRequest struct {
	Username string `schema:"username,required"`
	Password string `schema:"password,required"`
}
