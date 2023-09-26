package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/muhrizqiardi/linkbox/internal/service"
)

type AuthMiddleware interface {
	OnlyAllowRegisteredUser(next http.Handler) http.Handler
}

type authMiddleware struct {
	lg *log.Logger
	as service.AuthService
	us service.UserService
}

func NewAuthMiddleware(lg *log.Logger, as service.AuthService, us service.UserService) *authMiddleware {
	return &authMiddleware{lg, as, us}
}

func (m *authMiddleware) OnlyAllowRegisteredUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			m.lg.Println("failed to get token cookie:", err)
			http.Redirect(w, r, "/log-in", http.StatusSeeOther)
			return
		}
		parsedToken, err := m.as.CheckIsValid(cookie.Value)
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
