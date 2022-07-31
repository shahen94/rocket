package render

import (
	"errors"
	"net/http"
	"strings"

	"github.com/CloudyKit/jet/v6"
)

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
	JetViews   *jet.Set
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Port            string
	ServiceName     string
	Secure          bool
}

func New(rnd, rootPath, port string, jetViews *jet.Set) *Render {
	renderer := Render{
		Renderer: rnd,
		RootPath: rootPath,
		Port:     port,
		JetViews: jetViews,
	}

	return &renderer
}

func (rnd *Render) Page(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {
	switch strings.ToLower(rnd.Renderer) {
	case "go":
		return rnd.GoPage(w, r, view, data)
	case "jet":
		return rnd.JetPage(w, r, view, variables, data)
	default:
		return errors.New("Unknown renderer")
	}
}
