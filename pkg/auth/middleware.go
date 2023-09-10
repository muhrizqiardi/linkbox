package auth

import (
	"context"
	"log"
	"net/http"

	"github.com/muhrizqiardi/linkbox/linkbox/pkg/user"
)

func AuthMiddleware(lg *log.Logger, as Service, us user.Service) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("token")
			if err != nil {
				lg.Println("failed to get token cookie:", err)
				http.Redirect(w, r, "/log-in", http.StatusSeeOther)
			}
			parsedToken, newToken, err := as.CheckIsValid(cookie.Value)
			if err != nil {
				lg.Println("failed to check token is valid:", err)
				http.Redirect(w, r, "/log-in", http.StatusSeeOther)
			}
			newCookie := http.Cookie{
				Name:   "token",
				Value:  newToken,
				MaxAge: 8 * 24 * 60 * 60,
			}
			http.SetCookie(w, &newCookie)

			foundUser, err := us.GetOneByID(parsedToken.UserID)
			if err != nil {
				lg.Println("user not found, user not authenticated:", err)
			}

			ctx := context.WithValue(r.Context(), "user", foundUser)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		})
	}
}
