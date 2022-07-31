package render

import (
	"fmt"
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

func (rnd *Render) JetPage(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	vars := jet.VarMap{}

	if variables == nil {
		vars = make(jet.VarMap)
	} else {
		vars = variables.(jet.VarMap)
	}

	td := &TemplateData{}

	if data != nil {
		td = data.(*TemplateData)
	}

	t, err := rnd.JetViews.GetTemplate(fmt.Sprintf("%s.jet", view))

	if err != nil {
		log.Println(err)
		return err
	}

	if err = t.Execute(w, vars, td); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
