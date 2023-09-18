package templates

import (
	"embed"
	"html/template"
	"io"

	"github.com/muhrizqiardi/linkbox/pkg/common"
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

func (t *Templates) RegisterPage(w io.Writer, data common.RegisterPageData) error {
	if err := t.tmpl.ExecuteTemplate(w, "pages-register.html", data); err != nil {
		return err
	}

	return nil
}

func (t *Templates) LogInPage(w io.Writer, data common.LogInPageData) error {
	if err := t.tmpl.ExecuteTemplate(w, "pages-log-in.html", data); err != nil {
		return err
	}

	return nil
}

func (t *Templates) IndexPage(w io.Writer, data common.IndexPageData) error {
	if err := t.tmpl.ExecuteTemplate(w, "pages-index.html", data); err != nil {
		return err
	}

	return nil
}

func (t *Templates) SearchPage(w io.Writer, data common.SearchPageData) error {
	if err := t.tmpl.ExecuteTemplate(w, "pages-search.html", data); err != nil {
		return err
	}

	return nil
}

func (t *Templates) SearchResultsFragment(w io.Writer, data common.SearchResultsFragmentData) error {
	if err := t.tmpl.ExecuteTemplate(w, "fragments-search-results.html", data); err != nil {
		return err
	}

	return nil
}

func (t *Templates) LinksInFolderPage(w io.Writer, data common.LinksInFolderPageData) error {
	if err := t.tmpl.ExecuteTemplate(w, "pages-links-in-folder.html", data); err != nil {
		return err
	}

	return nil
}

func (t *Templates) LinkFragment(w io.Writer, data common.LinkFragmentData) error {
	if err := t.tmpl.ExecuteTemplate(w, "fragments-link.html", data); err != nil {
		return err
	}

	return nil
}

func (t *Templates) EditLinkModalFragment(w io.Writer, data common.EditLinkModalFragmentData) error {
	if err := t.tmpl.ExecuteTemplate(w, "fragments-edit-link-modal.html", data); err != nil {
		return err
	}

	return nil
}

func (t *Templates) DeleteLinkConfirmationModalFragment(w io.Writer, data common.DeleteLinkConfirmationModalFragmentData) error {
	if err := t.tmpl.ExecuteTemplate(w, "fragments-delete-link-confirmation-modal.html", data); err != nil {
		return err
	}

	return nil
}
