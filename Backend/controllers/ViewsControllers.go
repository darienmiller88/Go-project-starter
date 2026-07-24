package controllers

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

type ViewsController struct {
	Router    *chi.Mux
	templates *template.Template
}

func NewViewsController() *ViewsController {
	base        := []string{"templates/Base.html"}
	partials, _ := filepath.Glob("./templates/partials/*.html")

	htmlFiles   := append(base, partials...)

	pages, _    := filepath.Glob("./templates/pages/*.html")
	htmlFiles    = append(htmlFiles, pages...)

	vc := &ViewsController{
		Router:    chi.NewRouter(),
		templates: template.Must(template.ParseFiles(htmlFiles...)),
	}

	vc.registerViewRoutes()

	return vc
}

func (v *ViewsController) registerViewRoutes() {
	v.Router.Get("/home", v.homePage)
}


func (v *ViewsController) homePage(response http.ResponseWriter, request *http.Request) {
	err := v.templates.ExecuteTemplate(response, "home", nil)

	if err != nil {
		http.Error(response, "Error rendering template", http.StatusInternalServerError)
	}
}