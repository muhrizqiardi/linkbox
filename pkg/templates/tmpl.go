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

type Templates struct {
	tmpl *template.Template
}

func NewTemplates() (*Templates, error) {
	tmpl, err := template.ParseFS(tmplFS, "*.html")
	if err != nil {
		return &Templates{}, err
	}

	return &Templates{tmpl}, nil
}

type MetaData struct {
	Title       string
	Description string
	ImageURL    string
}

type RegisterPageData struct {
	Errors []string
	MetaData
}

func (t *Templates) RegisterPage(w io.Writer, data RegisterPageData) error {
	if err := t.tmpl.ExecuteTemplate(w, "pages-register.html", data); err != nil {
		return err
	}

	return nil
}

type LogInPageData struct {
	Errors []string
	MetaData
}

func (t *Templates) LogInPage(w io.Writer, data LogInPageData) error {
	if err := t.tmpl.ExecuteTemplate(w, "pages-log-in.html", data); err != nil {
		return err
	}

	return nil
}

type IndexPageData struct {
	User    user.UserEntity
	Folders []folder.FolderEntity
	Links   []link.LinkEntity
	MetaData
}

func (t *Templates) IndexPage(w io.Writer, data IndexPageData) error {
	if err := t.tmpl.ExecuteTemplate(w, "pages-index.html", data); err != nil {
		return err
	}

	return nil
}

type LinksInFolderPageData struct {
	User    user.UserEntity
	Folder  folder.FolderEntity
	Folders []folder.FolderEntity
	Links   []link.LinkEntity
	MetaData
}

func (t *Templates) LinksInFolderPage(w io.Writer, data LinksInFolderPageData) error {
	if err := t.tmpl.ExecuteTemplate(w, "pages-links-in-folder.html", data); err != nil {
		return err
	}

	return nil
}

type LinkFragmentData struct {
	Link link.LinkEntity
}

func (t *Templates) LinkFragment(w io.Writer, data struct{ Link link.LinkEntity }) error {
	if err := t.tmpl.ExecuteTemplate(w, "fragments-link.html", data); err != nil {
		return err
	}

	return nil
}

type EditLinkModalFragmentData struct {
	User    user.UserEntity
	Folders []folder.FolderEntity
	Link    link.LinkEntity
}

func (t *Templates) EditLinkModalFragment(w io.Writer, data EditLinkModalFragmentData) error {
	if err := t.tmpl.ExecuteTemplate(w, "fragments-edit-link-modal.html", data); err != nil {
		return err
	}

	return nil
}
