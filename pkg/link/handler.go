package link

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/gorilla/schema"
)

type templater interface {
	LinkFragment(w io.Writer, data struct{ Link LinkEntity }) error
}

type Handler struct {
	lg *log.Logger
	ls Service
	t  templater
}

func NewHandler(lg *log.Logger, ls Service, t templater) *Handler {
	return &Handler{lg, ls, t}
}

func (h *Handler) HandleCreateLink(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if err := r.ParseForm(); err != nil {
		h.lg.Println("failed to parse form body:", err)
		http.Error(w, "Failed to parse form body", http.StatusBadRequest)
	}

	var payload CreateLinkDTO
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

func (h *Handler) HandleUpdateLink(w http.ResponseWriter, r *http.Request) {
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

	var payload UpdateLinkDTO
	if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
		h.lg.Println("failed to parse form body:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	l, err := h.ls.UpdateOneByID(linkID, payload)
	if err != nil {
		h.lg.Println("failed to create link:", err)
	}

	// http.Redirect(w, r, fmt.Sprintf("/links/%d/card", l.ID), http.StatusSeeOther)
	if err := h.t.LinkFragment(w, struct{ Link LinkEntity }{l}); err != nil {
		h.lg.Println("failed to execute fragment template:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	return
}
