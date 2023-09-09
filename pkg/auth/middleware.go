package auth

import (
	"context"
	"log"
	"net/http"
)

func AuthMiddleware(lg *log.Logger, as Service) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("token")
			if err != nil {
				http.Redirect(w, r, "/log-in", http.StatusSeeOther)
			}
			parsedToken, newToken, err := as.CheckIsValid(cookie.Value)
			if err != nil {
				http.Redirect(w, r, "/log-in", http.StatusSeeOther)
			}
			newCookie := http.Cookie{
				Name:   "token",
				Value:  newToken,
				MaxAge: 8 * 24 * 60 * 60,
			}
			http.SetCookie(w, &newCookie)
			ctx := context.WithValue(r.Context(), "userId", parsedToken.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
			return
		})
	}
}