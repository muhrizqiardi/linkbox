package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/muhrizqiardi/linkbox/internal/entities/request"
	"github.com/muhrizqiardi/linkbox/internal/service"
)

type FolderHandler interface {
	HandleCreateFolder(w http.ResponseWriter, r *http.Request)
}

type folderHandler struct {
	lg *log.Logger
	fs service.FolderService
}

func NewFolderHandler(lg *log.Logger, fs service.FolderService) *folderHandler {
	return &folderHandler{lg, fs}
}

func (h *folderHandler) HandleCreateFolder(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.lg.Println("failed to parse form body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var payload request.CreateFolderRequest
	if err := schema.NewDecoder().Decode(&payload, r.PostForm); err != nil {
		h.lg.Println("failed to decode form body into a struct:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	f, err := h.fs.Create(payload)
	if err != nil {
		h.lg.Println("failed to create folder:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/folders/%d/links", f.ID), http.StatusSeeOther)
	return
}
