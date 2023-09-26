package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
	"github.com/muhrizqiardi/linkbox/internal/entities"
	"github.com/muhrizqiardi/linkbox/internal/entities/request"
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
}

func NewLinkHandler(lg *log.Logger, ls service.LinkService, t template.Executor) *linkHandler {
	return &linkHandler{lg, ls, t}
}

func (h *linkHandler) HandleCreateLink(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if err := r.ParseForm(); err != nil {
		h.lg.Println("failed to parse form body:", err)
		http.Error(w, "Failed to parse form body", http.StatusBadRequest)
	}

	var payload request.CreateLinkRequest
	if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
		h.lg.Println("failed to parse form body:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	l, err := h.ls.Create(payload)
	if err != nil {
		h.lg.Println("failed to create link:", err)
	}
	// TODO: use HTMX
	http.Redirect(w, r, fmt.Sprintf("/folders/%d/links#%d", l.FolderID, l.ID), http.StatusSeeOther)
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
	linkID, err := strconv.Atoi(chi.URLParam(r, "linkID"))
	if err != nil {
		h.lg.Println("failed to parse link ID from URL:", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	if err := r.ParseForm(); err != nil {
		h.lg.Println("failed to parse form body:", err)
		http.Error(w, "Failed to parse form body", http.StatusBadRequest)
	}

	var payload request.UpdateLinkRequest
	if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
		h.lg.Println("failed to parse form body:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	l, err := h.ls.UpdateOneByID(linkID, payload)
	if err != nil {
		h.lg.Println("failed to create link:", err)
	}

	if err := h.tx.LinkFragment(w, entities.LinkFragmentData{Link: l}); err != nil {
		h.lg.Println("failed to execute fragment template:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
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
