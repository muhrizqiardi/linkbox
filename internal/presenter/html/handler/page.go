package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/muhrizqiardi/linkbox/internal/constant"
	"github.com/muhrizqiardi/linkbox/internal/entities"
	"github.com/muhrizqiardi/linkbox/internal/entities/request"
	"github.com/muhrizqiardi/linkbox/internal/model"
	"github.com/muhrizqiardi/linkbox/internal/presenter/html/template"
	"github.com/muhrizqiardi/linkbox/internal/service"
	// "github.com/muhrizqiardi/linkbox/pkg/folder"
)

type PageHandler interface {
	HandleIndexPage(w http.ResponseWriter, r *http.Request)
	HandleSearchPage(w http.ResponseWriter, r *http.Request)
	HandleNewFolderModalFragment(w http.ResponseWriter, r *http.Request)
	HandleLinksInFolderPage(w http.ResponseWriter, r *http.Request)
	HandleLinksFragment(w http.ResponseWriter, r *http.Request)
	HandleNewLinkModalFragment(w http.ResponseWriter, r *http.Request)
	HandleEditLinkModalFragment(w http.ResponseWriter, r *http.Request)
	HandleRegisterPage(w http.ResponseWriter, r *http.Request)
	HandleLogInPage(w http.ResponseWriter, r *http.Request)
}

type pageHandler struct {
	lg *log.Logger
	fs service.FolderService
	ls service.LinkService
	as service.AuthService
	tx template.Executor
}

func NewPageHandler(
	lg *log.Logger,
	fs service.FolderService,
	ls service.LinkService,
	as service.AuthService,
	tx template.Executor,
) *pageHandler {
	return &pageHandler{lg, fs, ls, as, tx}
}

func (h *pageHandler) HandleIndexPage(w http.ResponseWriter, r *http.Request) {
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
	foundUser, ok := uCtx.(model.UserModel)
	if !ok {
		h.lg.Println("failed to get user data passed from middleware")
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	folders, err := h.fs.GetMany(foundUser.ID, request.GetManyFoldersRequest{
		Sort:    constant.GetManyFoldersSortDESC,
		OrderBy: constant.GetManyFoldersOrderByUpdatedAt,
		Limit:   20,
		Offset:  0,
	})
	if err != nil {
		h.lg.Println("failed to find folders related to user", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	df, _ := h.fs.GetOneByUniqueName("default", foundUser.ID)
	links, err := h.ls.GetManyInsideDefaultFolder(foundUser.ID, request.GetManyLinksInsideFolderRequest{
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

	if err := h.tx.IndexPage(w, entities.IndexPageData{
		User:         foundUser,
		Folders:      folders,
		Links:        links,
		FolderID:     df.ID,
		NextPage:     page + 1,
		PageMetaData: entities.PageMetaData{Title: "Home - Linkbox", Description: "Home of the Linkbox app", ImageURL: ""},
	}); err != nil {
		h.lg.Println("failed to render page:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *pageHandler) HandleSearchPage(w http.ResponseWriter, r *http.Request) {
	uCtx := r.Context().Value("user")
	foundUser, ok := uCtx.(model.UserModel)
	if !ok {
		h.lg.Println("failed to get user data passed from middleware")
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	ff, err := h.fs.GetMany(foundUser.ID, request.GetManyFoldersRequest{
		Sort:    constant.GetManyFoldersSortDESC,
		OrderBy: constant.GetManyFoldersOrderByUpdatedAt,
		Limit:   20,
		Offset:  0,
	})
	if err != nil {
		h.lg.Println("failed to find folders related to user", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	if err := h.tx.SearchPage(w, entities.SearchPageData{
		User:    foundUser,
		Folders: ff,
		Links:   []model.LinkModel{},
		PageMetaData: entities.PageMetaData{
			Title:       "Search - Linkbox",
			Description: "Search for links",
			ImageURL:    "",
		},
	}); err != nil {
		h.lg.Println("failed to render page:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

func (h *pageHandler) HandleNewFolderModalFragment(w http.ResponseWriter, r *http.Request) {
	uCtx := r.Context().Value("user")
	u, _ := uCtx.(model.UserModel)

	h.tx.NewFolderModalFragment(w, entities.NewFolderModalFragmentData{User: u})
	return

}

func (h *pageHandler) HandleLinksInFolderPage(w http.ResponseWriter, r *http.Request) {
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
	foundUser, ok := uCtx.(model.UserModel)
	if !ok {
		h.lg.Println("failed to get user data passed from middleware")
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	folders, err := h.fs.GetMany(foundUser.ID, request.GetManyFoldersRequest{
		Sort:    constant.GetManyFoldersSortDESC,
		OrderBy: constant.GetManyFoldersOrderByUpdatedAt,
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
	links, err := h.ls.GetManyInsideFolder(foundUser.ID, folder.ID, request.GetManyLinksInsideFolderRequest{
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

	if err := h.tx.LinksInFolderPage(w, entities.LinksInFolderPageData{
		User:     foundUser,
		Folder:   folder,
		Folders:  folders,
		Links:    links,
		FolderID: folder.ID,
		NextPage: page + 1,
		PageMetaData: entities.PageMetaData{
			Title:       "Folders: " + folder.UniqueName + " - Linkbox",
			Description: "Home of the Linkbox app",
			ImageURL:    "",
		},
	}); err != nil {
		h.lg.Println("failed to render page:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *pageHandler) HandleLinksFragment(w http.ResponseWriter, r *http.Request) {
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
	u, _ := uCtx.(model.UserModel)
	ll, err := h.ls.GetManyInsideFolder(u.ID, folderID, request.GetManyLinksInsideFolderRequest{
		Limit:   itemPerPage,
		Offset:  (page - 1) * itemPerPage,
		OrderBy: orderBy,
		Sort:    sort,
	})
	if err != nil {
		h.lg.Println("failed to fetch links inside folder:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if len(ll) > 0 {
		h.tx.LinksFragment(w, entities.LinksFragmentData{
			Links:    ll,
			FolderID: folderID,
			NextPage: page + 1,
		})
	}
	return
}

func (h *pageHandler) HandleNewLinkModalFragment(w http.ResponseWriter, r *http.Request) {
	uCtx := r.Context().Value("user")
	u, _ := uCtx.(model.UserModel)
	ff, _ := h.fs.GetMany(u.ID, request.GetManyFoldersRequest{
		Sort:    constant.GetManyFoldersSortDESC,
		OrderBy: constant.GetManyFoldersOrderByUpdatedAt,
		Limit:   100,
		Offset:  0,
	})

	h.tx.NewLinkModalFragment(w, entities.NewLinkModalFragmentData{
		User:             u,
		Folders:          ff,
		InitialFormValue: request.CreateLinkRequest{},
		Errors:           []string{},
	})
	return
}

func (h *pageHandler) HandleEditLinkModalFragment(w http.ResponseWriter, r *http.Request) {
	linkID, err := strconv.Atoi(chi.URLParam(r, "linkID"))
	if err != nil {
		h.lg.Println("failed to parse link ID from URL:", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	l, err := h.ls.GetOneByID(linkID)
	if err != nil {
		h.lg.Println("failed to get link data")
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	uCtx := r.Context().Value("user")
	foundUser, ok := uCtx.(model.UserModel)
	if !ok {
		h.lg.Println("failed to get user data passed from middleware")
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	folders, err := h.fs.GetMany(foundUser.ID, request.GetManyFoldersRequest{
		Sort:    constant.GetManyFoldersSortDESC,
		OrderBy: constant.GetManyFoldersOrderByUpdatedAt,
		Limit:   20,
		Offset:  0,
	})

	if err != nil {
		h.lg.Println("failed to find folders related to user", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	if err := h.tx.EditLinkModalFragment(w, entities.EditLinkModalFragmentData{
		User: foundUser, Folders: folders, Link: l,
	}); err != nil {
		h.lg.Println("failed to render HTML fragment:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *pageHandler) HandleRegisterPage(w http.ResponseWriter, r *http.Request) {
	existingCookie, err := r.Cookie("token")
	if err == nil {
		_, err := h.as.CheckIsValid(existingCookie.Value)
		if err == nil {
			h.lg.Println("user already authenticated, redirecting")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	if err := h.tx.RegisterPage(w, entities.RegisterPageData{
		PageMetaData: entities.PageMetaData{
			Title:       "Register Account - Linkbox",
			Description: "Register a new account on Linkbox",
			ImageURL:    "",
		},
	}); err != nil {
		h.lg.Println("failed to render page:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *pageHandler) HandleLogInPage(w http.ResponseWriter, r *http.Request) {
	existingCookie, err := r.Cookie("token")
	if err == nil {
		_, err := h.as.CheckIsValid(existingCookie.Value)
		if err == nil {
			h.lg.Println("user already authenticated, redirecting")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	if err := h.tx.LogInPage(w, entities.LogInPageData{
		PageMetaData: entities.PageMetaData{
			Title:       "Register Account - Linkbox",
			Description: "Register a new account on Linkbox",
			ImageURL:    "",
		},
	}); err != nil {
		h.lg.Println("failed to render page:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
