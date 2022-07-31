package render

import (
	"fmt"
	"net/http"
	"text/template"
)

func (rnd *Render) GoPage(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {
	path := fmt.Sprintf("%s/views/%s.page.tmpl", rnd.RootPath, view)
	tmpl, err := template.ParseFiles(path)

	if err != nil {
		return err
	}

	td := &TemplateData{}

	if data != nil {
		td = data.(*TemplateData)
	}
	err = tmpl.Execute(w, td)

	if err != nil {
		return err
	}

	return nil
}
