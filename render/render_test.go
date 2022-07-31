package render

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var pageData = []struct {
	name          string
	renderer      string
	template      string
	errorExpected bool
	errorMessage  string
}{
	{"go_page", "go", "home", false, "error rendering go template"},
	{"go_page_no_template", "go", "no-template", true, "No error rendering non-existent go template"},
	{"jet_page", "jet", "home", false, "error rendering jet template"},
	{"jet_page_no_template", "jet", "no-template", true, "No error rendering non-existent jet template"},
	{"invalid_renderer_engine", "invalid-renderer", "home", true, "No error when using invalid renderer engine"},
}

func TestRender_Page(t *testing.T) {

	for _, e := range pageData {
		r, err := http.NewRequest("GET", "/", nil)

		if err != nil {
			t.Errorf("Error creating request: %v", err)
		}

		w := httptest.NewRecorder()

		testRenderer.Renderer = e.renderer
		testRenderer.RootPath = "./test_data"

		err = testRenderer.Page(w, r, e.template, nil, nil)

		if e.errorExpected {
			if err == nil {
				t.Errorf("%s: %s", e.name, e.errorMessage)
			}
		} else {
			if err != nil {
				t.Errorf("%s: %s: %s", e.name, e.errorMessage, err.Error())
			}
		}
	}
}

func TestRender_GoPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}

	testRenderer.Renderer = "go"
	testRenderer.RootPath = "./test_data"

	err = testRenderer.Page(w, r, "home", nil, nil)

	if err != nil {
		t.Errorf("Error rendering page: %v", err)
	}
}

func TestRender_JetPage(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}

	testRenderer.Renderer = "jet"
	testRenderer.RootPath = "./test_data"

	err = testRenderer.Page(w, r, "home", nil, nil)

	if err != nil {
		t.Errorf("Error rendering page: %v", err)
	}
}

func TestRender_UnknownRenderer(t *testing.T) {
	w := httptest.NewRecorder()
	r, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		t.Errorf("Error creating request: %v", err)
	}

	testRenderer.Renderer = "unknown"
	testRenderer.RootPath = "./test_data"

	err = testRenderer.Page(w, r, "home", nil, nil)

	if err == nil {
		t.Errorf("Error rendering page: %v", err)
	}
}
