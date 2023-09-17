package common

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/auth"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/folder"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/link"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/page"
)

func Route(
	lg *log.Logger,
	ph page.Handler,
	ah auth.Handler,
	am auth.Middleware,
	lh link.Handler,
	fh folder.Handler,
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
