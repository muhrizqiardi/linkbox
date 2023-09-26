package route

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/muhrizqiardi/linkbox/internal/presenter/html/handler"
	"github.com/muhrizqiardi/linkbox/internal/presenter/html/middleware"
)

func DefineRoute(
	lg *log.Logger,
	ph handler.PageHandler,
	ah handler.AuthHandler,
	am middleware.AuthMiddleware,
	lh handler.LinkHandler,
	fh handler.FolderHandler,
) chi.Router {
	r := chi.NewRouter()

	r.Use(
		chiMiddleware.Logger,
	)

	r.Handle("/dist/*", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist"))))
	r.Handle("/node_modules/*", http.StripPrefix("/node_modules/", http.FileServer(http.Dir("./node_modules"))))

	r.Get("/register", ph.HandleRegisterPage)
	r.Post("/register", ah.HandleCreateUserAndLogIn)

	r.Get("/log-in", ph.HandleLogInPage)
	r.Post("/log-in", ah.HandleAuthLogIn)

	r.Get("/auth/delete", ah.HandleLogOut)

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
		r.Get("/search", ph.HandleSearchPage)
		r.Post("/search", lh.HandleSearch)
	})

	return r
}
