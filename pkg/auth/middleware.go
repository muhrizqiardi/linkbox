package auth

import (
	"context"
	"log"
	"net/http"

	"github.com/muhrizqiardi/linkbox/linkbox/pkg/user"
)

type Middleware struct {
	lg *log.Logger
	as Service
	us user.Service
}

func NewMiddleware(lg *log.Logger, as Service, us user.Service) *Middleware {
	return &Middleware{lg, as, us}
}

func (m *Middleware) OnlyAllowRegisteredUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			m.lg.Println("failed to get token cookie:", err)
			http.Redirect(w, r, "/log-in", http.StatusSeeOther)
			return
		}
		parsedToken, newToken, err := m.as.CheckIsValid(cookie.Value)
		if err != nil {
			m.lg.Println("failed to check token is valid:", err)
			c := &http.Cookie{
				Name:   "token",
				Value:  "",
				Path:   "/",
				MaxAge: -1,
			}
			http.SetCookie(w, c)
			http.Redirect(w, r, "/log-in", http.StatusSeeOther)
			return
		}
		newCookie := http.Cookie{
			Name:   "token",
			Value:  newToken,
			MaxAge: 8 * 24 * 60 * 60,
		}
		http.SetCookie(w, &newCookie)

		foundUser, err := m.us.GetOneByID(parsedToken.UserID)
		if err != nil {
			m.lg.Println("user not found, user not authenticated:", err)
			c := &http.Cookie{
				Name:   "token",
				Value:  "",
				Path:   "/",
				MaxAge: -1,
			}
			http.SetCookie(w, c)
			http.Redirect(w, r, "/log-in", http.StatusSeeOther)
			return
		}

		ctx := context.WithValue(r.Context(), "user", foundUser)
		next.ServeHTTP(w, r.WithContext(ctx))
		return
	})
}
