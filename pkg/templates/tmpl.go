package templates

import (
	"embed"
	"html/template"
	"io"

	"github.com/muhrizqiardi/linkbox/linkbox/pkg/folder"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/link"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/user"
)

//go:embed *.html
var tmplFS embed.FS

type RegisterPageData struct {
	Errors []string
}

func RegisterPage(w io.Writer, data RegisterPageData) error {
	tmpl, err := template.ParseFS(tmplFS, "pages-register.html")
	if err != nil {
		return err
	}

	if err := tmpl.ExecuteTemplate(w, "pages-register.html", data); err != nil {
		return err
	}

	return nil
}

type LogInPageData struct {
	Errors []string
}

func LogInPage(w io.Writer, data LogInPageData) error {
	tmpl, err := template.ParseFS(tmplFS, "pages-log-in.html")
	if err != nil {
		return err
	}

	if err := tmpl.ExecuteTemplate(w, "pages-log-in.html", data); err != nil {
		return err
	}

	return nil
}

type IndexPageData struct {
	User    user.UserEntity
	Folders []folder.FolderEntity
	Links   []link.LinkEntity
}

func IndexPage(w io.Writer, data IndexPageData) error {
	tmpl, err := template.ParseFS(tmplFS, "pages-index.html", "partials-head.html", "partials-new-link-modal.html", "partials-sidebar.html", "partials-index-page-navbar.html", "partials-new-folder-modal.html")
	if err != nil {
		return err
	}

	if err := tmpl.ExecuteTemplate(w, "pages-index.html", data); err != nil {
		return err
	}

	return nil
}

type LinksInFolderPageData struct {
	User    user.UserEntity
	Folder  folder.FolderEntity
	Folders []folder.FolderEntity
	Links   []link.LinkEntity
}

func LinksInFolderPage(w io.Writer, data LinksInFolderPageData) error {
	tmpl, err := template.ParseFS(tmplFS, "pages-links-in-folder.html", "partials-head.html", "partials-new-link-modal.html", "partials-sidebar.html", "partials-navbar.html", "partials-new-folder-modal.html")
	if err != nil {
		return err
	}

	if err := tmpl.ExecuteTemplate(w, "pages-links-in-folder.html", data); err != nil {
		return err
	}

	return nil
}
