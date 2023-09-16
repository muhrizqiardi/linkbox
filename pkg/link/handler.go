package link

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

type Handler struct {
	lg *log.Logger
	ls Service
}

func NewHandler(lg *log.Logger, ls Service) *Handler {
	return &Handler{lg, ls}
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
	http.Redirect(w, r, fmt.Sprintf("/folders/%d/links#%d", l.FolderID, l.ID), http.StatusSeeOther)
	return
}
