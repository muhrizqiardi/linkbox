package templates

import (
	"embed"
	"html/template"
	"io"

	"github.com/muhrizqiardi/linkbox/linkbox/pkg/folder"
	"github.com/muhrizqiardi/linkbox/linkbox/pkg/user"
)

//go:embed *.html
var tmplFS embed.FS

type RegisterPageData struct {
	Errors []string
}

func RegisterPage(w io.Writer, data RegisterPageData) error {
	tmpl, err := template.ParseFS(tmplFS, "register.html")
	if err != nil {
		return err
	}

	if err := tmpl.ExecuteTemplate(w, "register.html", data); err != nil {
		return err
	}

	return nil
}

type LogInPageData struct {
	Errors []string
}

func LogInPage(w io.Writer, data LogInPageData) error {
	tmpl, err := template.ParseFS(tmplFS, "log-in.html")
	if err != nil {
		return err
	}

	if err := tmpl.ExecuteTemplate(w, "log-in.html", data); err != nil {
		return err
	}

	return nil
}

type IndexPageData struct {
	User    user.UserEntity
	Folders []folder.FolderEntity
}

func IndexPage(w io.Writer, data IndexPageData) error {
	tmpl, err := template.ParseFS(tmplFS, "index.html")
	if err != nil {
		return err
	}

	if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
		return err
	}

	return nil
}
