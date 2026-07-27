package controllers

import (
	"GoProjectStarter/Backend/services"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

type ViewsController struct {
	Router      *chi.Mux
	templates   *template.Template
	userService services.UserService
}

func NewViewsController(userService services.UserService) *ViewsController {
	base := []string{"templates/Base.html"}
	partials, _ := filepath.Glob("./templates/partials/*.html")

	htmlFiles := append(base, partials...)

	pages, _ := filepath.Glob("./templates/pages/*.html")
	htmlFiles = append(htmlFiles, pages...)

	vc := &ViewsController{
		Router:      chi.NewRouter(),
		templates:   template.Must(template.ParseFiles(htmlFiles...)),
		userService: userService,
	}

	vc.registerViewRoutes()

	return vc
}

func (v *ViewsController) registerViewRoutes() {
	v.Router.Get("/home", v.homePage)
}

func (v *ViewsController) homePage(response http.ResponseWriter, request *http.Request) {
	usersResult := v.userService.GetUsers()

	if usersResult.Err != nil {
		http.Error(response, usersResult.Err.Error(), usersResult.StatusCode)
		return
	}

	users := usersResult.ResultData
	data := map[string]any{
		"Users": users,
	}

	err := v.templates.ExecuteTemplate(response, "home", data)

	if err != nil {
		http.Error(response, "Error rendering template", http.StatusInternalServerError)
	}
}
