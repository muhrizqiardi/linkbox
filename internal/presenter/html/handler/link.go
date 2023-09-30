package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
	"github.com/lib/pq"
	"github.com/muhrizqiardi/linkbox/internal/constant"
	"github.com/muhrizqiardi/linkbox/internal/entities"
	"github.com/muhrizqiardi/linkbox/internal/entities/request"
	"github.com/muhrizqiardi/linkbox/internal/entities/response"
	"github.com/muhrizqiardi/linkbox/internal/model"
	"github.com/muhrizqiardi/linkbox/internal/presenter/html/template"
	"github.com/muhrizqiardi/linkbox/internal/service"
)

type LinkHandler interface {
	HandleCreateLink(w http.ResponseWriter, r *http.Request)
	HandleSearch(w http.ResponseWriter, r *http.Request)
	HandleUpdateLink(w http.ResponseWriter, r *http.Request)
	HandleDeleteLink(w http.ResponseWriter, r *http.Request)
	HandleDeleteLinkConfirmationModal(w http.ResponseWriter, r *http.Request)
}

type linkHandler struct {
	lg *log.Logger
	ls service.LinkService
	tx template.Executor
	fs service.FolderService
}

func NewLinkHandler(lg *log.Logger, ls service.LinkService, t template.Executor, fs service.FolderService) *linkHandler {
	return &linkHandler{lg, ls, t, fs}
}

func (h *linkHandler) HandleCreateLink(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	uCtx := r.Context().Value("user")
	u, _ := uCtx.(model.UserModel)

	if err := r.ParseForm(); err != nil {
		h.lg.Println("failed to parse form body:", err)
		ff, _ := h.fs.GetMany(u.ID, request.GetManyFoldersRequest{
			OrderBy: constant.GetManyFoldersOrderByUpdatedAt,
			Sort:    constant.GetManyFoldersSortDESC,
			Limit:   100,
			Offset:  0,
		})
		w.Header().Set("HX-Retarget", "#new_link_modal")
		w.Header().Set("HX-Reswap", "outerHTML")
		h.tx.NewLinkModalFragment(w, entities.NewLinkModalFragmentData{
			User:    u,
			Folders: ff,
			Errors:  []string{constant.ErrCreateLink.Error()},
		})
		return
	}

	var payload request.CreateLinkRequest
	if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
		h.lg.Println("failed to parse form body:", err)
		ff, _ := h.fs.GetMany(u.ID, request.GetManyFoldersRequest{
			OrderBy: constant.GetManyFoldersOrderByUpdatedAt,
			Sort:    constant.GetManyFoldersSortDESC,
			Limit:   100,
			Offset:  0,
		})
		w.Header().Set("HX-Retarget", "#new_link_modal")
		w.Header().Set("HX-Reswap", "outerHTML")
		h.tx.NewLinkModalFragment(w, entities.NewLinkModalFragmentData{
			User:    u,
			Folders: ff,
			Errors:  []string{constant.ErrCreateLink.Error()},
		})
		return
	}
	l, err := h.ls.Create(payload)
	if err != nil {
		h.lg.Println("failed to create link:", err)
		w.Header().Set("HX-Retarget", "#new_link_modal")
		w.Header().Set("HX-Reswap", "outerHTML")
		ff, _ := h.fs.GetMany(u.ID, request.GetManyFoldersRequest{
			OrderBy: constant.GetManyFoldersOrderByUpdatedAt,
			Sort:    constant.GetManyFoldersSortDESC,
			Limit:   100,
			Offset:  0,
		})
		if err, ok := err.(*pq.Error); ok {
			switch string(err.Code) {
			case "23505":
				h.tx.NewLinkModalFragment(w, entities.NewLinkModalFragmentData{
					User:             u,
					Folders:          ff,
					InitialFormValue: payload,
					Errors:           []string{constant.ErrDuplicateURL.Error()},
				})
				return
			}
		}

		h.tx.NewLinkModalFragment(w, entities.NewLinkModalFragmentData{
			User:             u,
			Folders:          ff,
			InitialFormValue: payload,
			Errors:           []string{constant.ErrCreateLink.Error()},
		})
	}

	redirectTo := fmt.Sprintf("/folders/%d/links#link_%d", l.FolderID, l.ID)
	w.Header().Set("HX-Redirect", redirectTo)
	http.Redirect(w, r, redirectTo, http.StatusSeeOther)
	return
}

func (h *linkHandler) HandleSearch(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.lg.Println("failed to parse form body:", err)
		http.Error(w, "Failed to parse form body", http.StatusBadRequest)
	}
	query := r.PostForm.Get("query")
	uCtx := r.Context().Value("user")
	foundUser, ok := uCtx.(model.UserModel)
	if !ok {
		h.lg.Println("failed to get user data passed from middleware")
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	ll, err := h.ls.SearchFullText(foundUser.ID, query)
	if err != nil {
		h.lg.Println("failed to search full-text for links:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if err := h.tx.SearchResultsFragment(w, entities.SearchResultsFragmentData{
		Links: ll,
	}); err != nil {
		h.lg.Println("failed to execute fragment template:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	return
}

func (h *linkHandler) HandleUpdateLink(w http.ResponseWriter, r *http.Request) {
	uCtx := r.Context().Value("user")
	u, _ := uCtx.(model.UserModel)

	linkID, err := strconv.Atoi(chi.URLParam(r, "linkID"))
	if err != nil {
		h.lg.Println("failed to parse link ID from URL:", err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	defer r.Body.Close()
	if err := r.ParseForm(); err != nil {
		h.lg.Println("failed to parse form body:", err)
		ff, _ := h.fs.GetMany(u.ID, request.GetManyFoldersRequest{
			OrderBy: constant.GetManyFoldersOrderByUpdatedAt,
			Sort:    constant.GetManyFoldersSortDESC,
			Limit:   100,
			Offset:  0,
		})
		w.Header().Set("HX-Retarget", "#new_link_modal")
		w.Header().Set("HX-Reswap", "outerHTML")
		h.tx.EditLinkModalFragment(w, entities.EditLinkModalFragmentData{
			User:    u,
			Folders: ff,
			Link:    model.LinkModel{},
		})
		return
	}

	var payload request.UpdateLinkRequest
	if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
		h.lg.Println("failed to parse form body:", err)
		ff, _ := h.fs.GetMany(u.ID, request.GetManyFoldersRequest{
			OrderBy: constant.GetManyFoldersOrderByUpdatedAt,
			Sort:    constant.GetManyFoldersSortDESC,
			Limit:   100,
			Offset:  0,
		})
		w.Header().Set("HX-Retarget", "#edit_link_modal")
		w.Header().Set("HX-Reswap", "outerHTML")
		h.tx.EditLinkModalFragment(w, entities.EditLinkModalFragmentData{
			User:    u,
			Folders: ff,
			Link:    model.LinkModel{},
		})
		return
	}
	l, err := h.ls.UpdateOneByID(linkID, payload)
	if err != nil {
		h.lg.Println("failed to update link:", err)
		w.Header().Set("HX-Retarget", "#edit_link_modal")
		w.Header().Set("HX-Reswap", "outerHTML")
		ff, _ := h.fs.GetMany(u.ID, request.GetManyFoldersRequest{
			OrderBy: constant.GetManyFoldersOrderByUpdatedAt,
			Sort:    constant.GetManyFoldersSortDESC,
			Limit:   100,
			Offset:  0,
		})
		if err, ok := err.(*pq.Error); ok {
			switch string(err.Code) {
			case "23505":
				h.tx.EditLinkModalFragment(w, entities.EditLinkModalFragmentData{
					User:    u,
					Folders: ff,
					Link: model.LinkModel{
						URL:         payload.URL,
						Title:       payload.Title,
						Description: payload.Description,
						FolderID:    payload.FolderID,
					},
					Errors: []string{constant.ErrDuplicateURL.Error()},
				})
				return
			}
		}

		h.tx.EditLinkModalFragment(w, entities.EditLinkModalFragmentData{
			User:    u,
			Folders: ff,
			Errors:  []string{constant.ErrUpdateLink.Error()},
		})
		return
	}

	w.Header().Set("HX-Trigger", "close-edit-link-modal")
	h.tx.LinkFragment(w, entities.LinkFragmentData{
		Link: response.LinkWithMediaResponse{
			ID:          l.ID,
			URL:         l.URL,
			Title:       l.Title,
			Description: l.Description,
			UserID:      l.UserID,
			FolderID:    l.FolderID,
			Media:       []response.LinkWithMediaResponseMedia{},
			CreatedAt:   l.CreatedAt,
			UpdatedAt:   l.UpdatedAt,
		},
	})
	return
}

func (h *linkHandler) HandleDeleteLinkConfirmationModal(w http.ResponseWriter, r *http.Request) {
	linkID, err := strconv.Atoi(chi.URLParam(r, "linkID"))
	if err != nil {
		h.lg.Println("failed to parse link ID from URL:", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := h.tx.DeleteLinkConfirmationModalFragment(
		w,
		entities.DeleteLinkConfirmationModalFragmentData{LinkID: linkID},
	); err != nil {
		h.lg.Println("failed to execute fragment template:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	return
}

func (h *linkHandler) HandleDeleteLink(w http.ResponseWriter, r *http.Request) {
	linkID, err := strconv.Atoi(chi.URLParam(r, "linkID"))
	if err != nil {
		h.lg.Println("failed to parse link ID from URL:", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if _, err := h.ls.DeleteOneByID(linkID); err != nil {
		h.lg.Println("failed to delete link ID:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	return
}
