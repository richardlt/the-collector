package server

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

// Template : template struct
type Template struct {
	Templates *template.Template
}

// Render : render template
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}
