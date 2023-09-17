package route

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/common"
)

func Route(
	lg *log.Logger,
	ph common.PageHandler,
	ah common.AuthHandler,
	am common.AuthMiddleware,
	lh common.LinkHandler,
	fh common.FolderHandler,
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
		r.Get("/links/{linkID}/delete", lh.HandleDeleteLinkConfirmationModal)
		r.Delete("/links/{linkID}", lh.HandleDeleteLink)
	})

	return r
}
