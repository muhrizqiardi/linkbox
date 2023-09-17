package link

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/common"
)

type handler struct {
	lg *log.Logger
	ls common.LinkService
	t  common.Templater
}

func NewHandler(lg *log.Logger, ls common.LinkService, t common.Templater) *handler {
	return &handler{lg, ls, t}
}

func (h *handler) HandleCreateLink(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if err := r.ParseForm(); err != nil {
		h.lg.Println("failed to parse form body:", err)
		http.Error(w, "Failed to parse form body", http.StatusBadRequest)
	}

	var payload common.CreateLinkDTO
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

func (h *handler) HandleUpdateLink(w http.ResponseWriter, r *http.Request) {
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

	var payload common.UpdateLinkDTO
	if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
		h.lg.Println("failed to parse form body:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	l, err := h.ls.UpdateOneByID(linkID, payload)
	if err != nil {
		h.lg.Println("failed to create link:", err)
	}

	if err := h.t.LinkFragment(w, common.LinkFragmentData{Link: l}); err != nil {
		h.lg.Println("failed to execute fragment template:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	return
}

func (h *handler) HandleDeleteLinkConfirmationModal(w http.ResponseWriter, r *http.Request) {
	linkID, err := strconv.Atoi(chi.URLParam(r, "linkID"))
	if err != nil {
		h.lg.Println("failed to parse link ID from URL:", err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	if err := h.t.DeleteLinkConfirmationModalFragment(
		w,
		common.DeleteLinkConfirmationModalFragmentData{LinkID: linkID},
	); err != nil {
		h.lg.Println("failed to execute fragment template:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	return
}

func (h *handler) HandleDeleteLink(w http.ResponseWriter, r *http.Request) {
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
