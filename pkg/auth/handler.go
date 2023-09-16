package auth

import (
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/user"
)

type Handler struct {
	lg *log.Logger
	as Service
	us user.Service
}

func NewHandler(lg *log.Logger, as Service, us user.Service) *Handler {
	return &Handler{lg, as, us}
}

func (h *Handler) HandleAuthLogIn(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.lg.Println("failed to parse form body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var payload LogInDTO
	if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
		h.lg.Println("failed to decode form body into a struct:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.as.LogIn(payload)
	if err != nil {
		h.lg.Println("failed to log in:", err)
		http.Error(w, "Failed to log in.", http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:   "token",
		Value:  token,
		MaxAge: 7 * 24 * 60 * 60,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

func (h *Handler) HandleCreateUserAndLogIn(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.lg.Println("failed to parse form body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var payload user.CreateUserDTO
	if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
		h.lg.Println("failed to decode form body into a struct:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := h.us.Create(payload)
	if err != nil {
		h.lg.Println("failed to create user:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	token, err := h.as.LogIn(LogInDTO{Username: user.Username, Password: user.Password})
	if err != nil {
		h.lg.Println("failed to log in:", err)
		http.Error(w, "Failed to log in. Account creation was success, so you can try logging in manually.", http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:   "token",
		Value:  token,
		MaxAge: 7 * 24 * 60 * 60,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return
}

func (h *Handler) HandleLogOut(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:   "token",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/log-in", http.StatusSeeOther)
	return
}
