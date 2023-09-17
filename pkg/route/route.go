package route

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type pageHandler interface {
	HandleIndexPage(w http.ResponseWriter, r *http.Request)
	HandleLinksInFolderPage(w http.ResponseWriter, r *http.Request)
	HandleEditLinkModalFragment(w http.ResponseWriter, r *http.Request)
	HandleRegisterPage(w http.ResponseWriter, r *http.Request)
	HandleLogInPage(w http.ResponseWriter, r *http.Request)
}

type authHandler interface {
	HandleAuthLogIn(w http.ResponseWriter, r *http.Request)
	HandleCreateUserAndLogIn(w http.ResponseWriter, r *http.Request)
	HandleLogOut(w http.ResponseWriter, r *http.Request)
}

type authMiddleware interface {
	OnlyAllowRegisteredUser(next http.Handler) http.Handler
}

type linkHandler interface {
	HandleCreateLink(w http.ResponseWriter, r *http.Request)
	HandleUpdateLink(w http.ResponseWriter, r *http.Request)
}

type folderHandler interface {
	HandleCreateFolder(w http.ResponseWriter, r *http.Request)
}

func Route(
	lg *log.Logger,
	ph pageHandler,
	ah authHandler,
	am authMiddleware,
	lh linkHandler,
	fh folderHandler,
) chi.Router {
	r := chi.NewRouter()

	r.Handle("/dist/*", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist"))))
	r.Handle("/node_modules/*", http.StripPrefix("/node_modules/", http.FileServer(http.Dir("./node_modules"))))

	r.Get("/register", ph.HandleRegisterPage)
	r.Post("/register", ah.HandleCreateUserAndLogIn)

	r.Get("/log-in", ph.HandleLogInPage)
	r.Post("/log-in", ah.HandleAuthLogIn)

	r.Get("/auth/delete", ah.HandleLogOut)

	// Needs Authentication
	r.Group(func(r chi.Router) {
		r.Use(
			am.OnlyAllowRegisteredUser,
		)

		r.Get("/", ph.HandleIndexPage)
		r.Post("/folders", fh.HandleCreateFolder)
		r.Get("/folders/{folderID}/links", ph.HandleLinksInFolderPage)
		r.Post("/links", lh.HandleCreateLink)
		r.Get("/links/{linkID}/edit", ph.HandleEditLinkModalFragment)
		r.Put("/links/{linkID}", lh.HandleUpdateLink)
	})

	return r
}
