package templates

import (
	"embed"
	"html/template"
	"io"
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
