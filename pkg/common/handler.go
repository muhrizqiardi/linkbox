package common

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/schema"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/auth"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/folder"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/link"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/templates"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/user"
)

func HandleIndexPage(lg *log.Logger, fs folder.Service, ls link.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
			lg.Println("failed to get user data passed from middleware")
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		folders, err := fs.GetMany(foundUser.ID, folder.GetManyFoldersDTO{
			Sort:    folder.GetManyFoldersSortDescending,
			OrderBy: folder.GetManyFoldersOrderByUpdatedAt,
			Limit:   20,
			Offset:  0,
		})
		if err != nil {
			lg.Println("failed to find folders related to user", err)
			http.Error(w, "", http.StatusNotFound)
			return
		}

		links, err := ls.GetManyInsideDefaultFolder(foundUser.ID, link.GetManyLinksInsideFolderDTO{
			Limit:   itemPerPage,
			Offset:  (page - 1) * itemPerPage,
			OrderBy: orderBy,
			Sort:    sort,
		})
		if err != nil {
			lg.Println("failed to find links inside default folder", err)
			http.Error(w, "", http.StatusNotFound)
			return
		}

		if err := templates.IndexPage(w, templates.IndexPageData{
			User:    foundUser,
			Folders: folders,
			Links:   links,
		}); err != nil {
			lg.Println("failed to render page:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func HandleLinksInFolderPage(lg *log.Logger, ls link.Service, fs folder.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
			lg.Println("failed to parse folder ID from URL:", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		uCtx := r.Context().Value("user")
		foundUser, ok := uCtx.(user.UserEntity)
		if !ok {
			lg.Println("failed to get user data passed from middleware")
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		folders, err := fs.GetMany(foundUser.ID, folder.GetManyFoldersDTO{
			Sort:    folder.GetManyFoldersSortDescending,
			OrderBy: folder.GetManyFoldersOrderByUpdatedAt,
			Limit:   20,
			Offset:  0,
		})
		if err != nil {
			lg.Println("failed to find folders related to user", err)
			http.Error(w, "", http.StatusNotFound)
			return
		}

		folder, err := fs.GetOneByID(folderID)
		if err != nil {
			lg.Println("failed to find folder detail", err)
			http.Error(w, "", http.StatusNotFound)
			return
		}
		links, err := ls.GetManyInsideFolder(foundUser.ID, folder.ID, link.GetManyLinksInsideFolderDTO{
			Limit:   itemPerPage,
			Offset:  (page - 1) * itemPerPage,
			OrderBy: orderBy,
			Sort:    sort,
		})
		if err != nil {
			lg.Println("failed to find links inside default folder", err)
			http.Error(w, "", http.StatusNotFound)
			return
		}

		if err := templates.LinksInFolderPage(w, templates.LinksInFolderPageData{
			User:    foundUser,
			Folder:  folder,
			Folders: folders,
			Links:   links,
		}); err != nil {
			lg.Println("failed to render page:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func HandleCreateLink(lg *log.Logger, ls link.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		if err := r.ParseForm(); err != nil {
			lg.Println("failed to parse form body:", err)
			http.Error(w, "Failed to parse form body", http.StatusBadRequest)
		}

		var payload link.CreateLinkDTO
		if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
			lg.Println("failed to parse form body:", err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		link, err := ls.Create(payload)
		if err != nil {
			lg.Println("failed to create link:", err)
		}
		http.Redirect(w, r, fmt.Sprintf("/folders/%d/links#%d", link.FolderID, link.ID), http.StatusSeeOther)
		return
	}
}

func HandleRegisterPage(lg *log.Logger, as auth.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		existingCookie, err := r.Cookie("token")
		if err == nil {
			_, newToken, err := as.CheckIsValid(existingCookie.Value)
			if err == nil {
				lg.Println("user already authenticated, redirecting")
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
			lg.Println("failed to render page:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func HandleCreateUser(lg *log.Logger, us user.Service, as auth.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			lg.Println("failed to parse form body:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var payload user.CreateUserDTO
		if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
			lg.Println("failed to decode form body into a struct:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		user, err := us.Create(payload)
		if err != nil {
			lg.Println("failed to create user:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		token, err := as.LogIn(auth.LogInDTO{Username: user.Username, Password: user.Password})
		if err != nil {
			lg.Println("failed to log in:", err)
			http.Error(w, "Failed to log in. Account creation was success, so you can try logging in manually.", http.StatusInternalServerError)
			return
		}

		cookie := http.Cookie{
			Name:   "token",
			Value:  token,
			MaxAge: 7 * 24 * 60 * 60,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}

func HandleCreateFolder(lg *log.Logger, fs folder.Service, as auth.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			lg.Println("failed to parse form body:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var payload folder.CreateFolderDTO
		if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
			lg.Println("failed to decode form body into a struct:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		f, err := fs.Create(payload)
		if err != nil {
			lg.Println("failed to create folder:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, fmt.Sprintf("/folders/%d/links", f.ID), http.StatusSeeOther)
		return
	}
}

func HandleLogInPage(lg *log.Logger, as auth.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		existingCookie, err := r.Cookie("token")
		if err == nil {
			_, newToken, err := as.CheckIsValid(existingCookie.Value)
			if err == nil {
				lg.Println("user already authenticated, redirecting")
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
			lg.Println("failed to render page:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func HandleAuthLogIn(lg *log.Logger, as auth.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			lg.Println("failed to parse form body:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var payload auth.LogInDTO
		if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
			lg.Println("failed to decode form body into a struct:", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		token, err := as.LogIn(payload)
		if err != nil {
			lg.Println("failed to log in:", err)
			http.Error(w, "Failed to log in.", http.StatusInternalServerError)
			return
		}

		cookie := http.Cookie{
			Name:   "token",
			Value:  token,
			MaxAge: 7 * 24 * 60 * 60,
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
}
