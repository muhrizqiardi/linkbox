package page

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/auth"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/folder"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/link"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/templates"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/user"
)

type Handler struct {
	lg *log.Logger
	fs folder.Service
	ls link.Service
	as auth.Service
}

func NewHandler(lg *log.Logger, fs folder.Service, ls link.Service, as auth.Service) *Handler {
	return &Handler{lg, fs, ls, as}
}

func (h *Handler) HandleIndexPage(w http.ResponseWriter, r *http.Request) {
	var page int = 1
	var itemPerPage int = 10
	var orderBy string = "updated_at"
	var sort string = "desc"
	if queryPage, _ := strconv.Atoi(r.URL.Query().Get("page")); queryPage != 0 {
		page = queryPage
	}
	if queryItemPerPage, _ := strconv.Atoi(r.URL.Query().Get("itemPerPage")); queryItemPerPage != 0 {
		itemPerPage = queryItemPerPage
	}
	if queryOrderBy := r.URL.Query().Get("orderBy"); queryOrderBy != "" {
		orderBy = queryOrderBy
	}
	if querySort := r.URL.Query().Get("sort"); querySort != "" {
		sort = querySort
	}

	uCtx := r.Context().Value("user")
	foundUser, ok := uCtx.(user.UserEntity)
	if !ok {
		h.lg.Println("failed to get user data passed from middleware")
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	folders, err := h.fs.GetMany(foundUser.ID, folder.GetManyFoldersDTO{
		Sort:    folder.GetManyFoldersSortDescending,
		OrderBy: folder.GetManyFoldersOrderByUpdatedAt,
		Limit:   20,
		Offset:  0,
	})
	if err != nil {
		h.lg.Println("failed to find folders related to user", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	links, err := h.ls.GetManyInsideDefaultFolder(foundUser.ID, link.GetManyLinksInsideFolderDTO{
		Limit:   itemPerPage,
		Offset:  (page - 1) * itemPerPage,
		OrderBy: orderBy,
		Sort:    sort,
	})
	if err != nil {
		h.lg.Println("failed to find links inside default folder", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	if err := templates.IndexPage(w, templates.IndexPageData{
		User:    foundUser,
		Folders: folders,
		Links:   links,
	}); err != nil {
		h.lg.Println("failed to render page:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func (h *Handler) HandleLinksInFolderPage(w http.ResponseWriter, r *http.Request) {
	var page int = 1
	var itemPerPage int = 10
	var orderBy string = "updated_at"
	var sort string = "desc"
	if queryPage, _ := strconv.Atoi(r.URL.Query().Get("page")); queryPage != 0 {
		page = queryPage
	}
	if queryItemPerPage, _ := strconv.Atoi(r.URL.Query().Get("itemPerPage")); queryItemPerPage != 0 {
		itemPerPage = queryItemPerPage
	}
	if queryOrderBy := r.URL.Query().Get("orderBy"); queryOrderBy != "" {
		orderBy = queryOrderBy
	}
	if querySort := r.URL.Query().Get("sort"); querySort != "" {
		sort = querySort
	}
	folderID, err := strconv.Atoi(chi.URLParam(r, "folderID"))
	if err != nil {
		h.lg.Println("failed to parse folder ID from URL:", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	uCtx := r.Context().Value("user")
	foundUser, ok := uCtx.(user.UserEntity)
	if !ok {
		h.lg.Println("failed to get user data passed from middleware")
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	folders, err := h.fs.GetMany(foundUser.ID, folder.GetManyFoldersDTO{
		Sort:    folder.GetManyFoldersSortDescending,
		OrderBy: folder.GetManyFoldersOrderByUpdatedAt,
		Limit:   20,
		Offset:  0,
	})
	if err != nil {
		h.lg.Println("failed to find folders related to user", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	folder, err := h.fs.GetOneByID(folderID)
	if err != nil {
		h.lg.Println("failed to find folder detail", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}
	links, err := h.ls.GetManyInsideFolder(foundUser.ID, folder.ID, link.GetManyLinksInsideFolderDTO{
		Limit:   itemPerPage,
		Offset:  (page - 1) * itemPerPage,
		OrderBy: orderBy,
		Sort:    sort,
	})
	if err != nil {
		h.lg.Println("failed to find links inside default folder", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	if err := templates.LinksInFolderPage(w, templates.LinksInFolderPageData{
		User:    foundUser,
		Folder:  folder,
		Folders: folders,
		Links:   links,
	}); err != nil {
		h.lg.Println("failed to render page:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) HandleRegisterPage(w http.ResponseWriter, r *http.Request) {
	existingCookie, err := r.Cookie("token")
	if err == nil {
		_, newToken, err := h.as.CheckIsValid(existingCookie.Value)
		if err == nil {
			h.lg.Println("user already authenticated, redirecting")
			http.SetCookie(w, &http.Cookie{
				Name:   "token",
				Value:  newToken,
				MaxAge: 8 * 24 * 60 * 60,
			})
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	if err := templates.RegisterPage(w, templates.RegisterPageData{}); err != nil {
		h.lg.Println("failed to render page:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) HandleLogInPage(w http.ResponseWriter, r *http.Request) {
	existingCookie, err := r.Cookie("token")
	if err == nil {
		_, newToken, err := h.as.CheckIsValid(existingCookie.Value)
		if err == nil {
			h.lg.Println("user already authenticated, redirecting")
			http.SetCookie(w, &http.Cookie{
				Name:   "token",
				Value:  newToken,
				MaxAge: 8 * 24 * 60 * 60,
			})
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	if err := templates.LogInPage(w, templates.LogInPageData{}); err != nil {
		h.lg.Println("failed to render page:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
