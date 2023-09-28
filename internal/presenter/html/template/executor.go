package template

import (
	"embed"
	goHTMLTemplate "html/template"
	"io"

	"github.com/muhrizqiardi/linkbox/internal/entities"
)

//go:embed templates/*.html
var tmplFS embed.FS

type Executor interface {
	RegisterPage(w io.Writer, data entities.RegisterPageData) error
	LogInPage(w io.Writer, data entities.LogInPageData) error
	SearchPage(w io.Writer, data entities.SearchPageData) error
	SearchResultsFragment(w io.Writer, data entities.SearchResultsFragmentData) error
	IndexPage(w io.Writer, data entities.IndexPageData) error
	LinksInFolderPage(w io.Writer, data entities.LinksInFolderPageData) error
	NewLinkModalFragment(w io.Writer, data entities.NewLinkModalFragmentData) error
	LinkFragment(w io.Writer, data entities.LinkFragmentData) error
	EditLinkModalFragment(w io.Writer, data entities.EditLinkModalFragmentData) error
	DeleteLinkConfirmationModalFragment(w io.Writer, data entities.DeleteLinkConfirmationModalFragmentData) error
}

type executor struct {
	tmpl *goHTMLTemplate.Template
}

func NewExecutor() (*executor, error) {
	tmpl, err := goHTMLTemplate.ParseFS(tmplFS, "templates/*.html")
	if err != nil {
		return &executor{}, err
	}

	return &executor{tmpl}, nil
}

func (t *executor) RegisterPage(w io.Writer, data entities.RegisterPageData) error {
	if err := t.tmpl.ExecuteTemplate(w, "pages-register.html", data); err != nil {
		return err
	}

	return nil
}

func (t *executor) LogInPage(w io.Writer, data entities.LogInPageData) error {
	if err := t.tmpl.ExecuteTemplate(w, "pages-log-in.html", data); err != nil {
		return err
	}

	return nil
}

func (t *executor) IndexPage(w io.Writer, data entities.IndexPageData) error {
	if err := t.tmpl.ExecuteTemplate(w, "pages-index.html", data); err != nil {
		return err
	}

	return nil
}

func (t *executor) SearchPage(w io.Writer, data entities.SearchPageData) error {
	if err := t.tmpl.ExecuteTemplate(w, "pages-search.html", data); err != nil {
		return err
	}

	return nil
}

func (t *executor) NewLinkModalFragment(w io.Writer, data entities.NewLinkModalFragmentData) error {
	if err := t.tmpl.ExecuteTemplate(w, "fragments-new-link-modal.html", data); err != nil {
		return err
	}

	return nil
}

func (t *executor) SearchResultsFragment(w io.Writer, data entities.SearchResultsFragmentData) error {
	if err := t.tmpl.ExecuteTemplate(w, "fragments-search-results.html", data); err != nil {
		return err
	}

	return nil
}

func (t *executor) LinksInFolderPage(w io.Writer, data entities.LinksInFolderPageData) error {
	if err := t.tmpl.ExecuteTemplate(w, "pages-links-in-folder.html", data); err != nil {
		return err
	}

	return nil
}

func (t *executor) LinkFragment(w io.Writer, data entities.LinkFragmentData) error {
	if err := t.tmpl.ExecuteTemplate(w, "fragments-link.html", data); err != nil {
		return err
	}

	return nil
}

func (t *executor) EditLinkModalFragment(w io.Writer, data entities.EditLinkModalFragmentData) error {
	if err := t.tmpl.ExecuteTemplate(w, "fragments-edit-link-modal.html", data); err != nil {
		return err
	}

	return nil
}

func (t *executor) DeleteLinkConfirmationModalFragment(w io.Writer, data entities.DeleteLinkConfirmationModalFragmentData) error {
	if err := t.tmpl.ExecuteTemplate(w, "fragments-delete-link-confirmation-modal.html", data); err != nil {
		return err
	}

	return nil
}
