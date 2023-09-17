package folder

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/muhrizqiardi/linkbox/pkg/common"
)

type handler struct {
	lg *log.Logger
	fs common.FolderService
}

func NewHandler(lg *log.Logger, fs common.FolderService) *handler {
	return &handler{lg, fs}
}

func (h *handler) HandleCreateFolder(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		h.lg.Println("failed to parse form body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var payload common.CreateFolderDTO
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
