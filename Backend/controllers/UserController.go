package controllers

import (
	"GoProjectStarter/Backend/services"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	Router      *chi.Mux
	templates   *template.Template
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {

	uc := &UserController{
		Router:      chi.NewRouter(),
		userService: userService,
		templates:   template.Must(template.ParseGlob("./templates/partials/*.html")),
	}

	uc.registerViewRoutes()

	return uc
}

func (u *UserController) registerViewRoutes() {
	u.Router.Post("/add-user", u.addUser)
	u.Router.Get("/get-users", u.getUsers)
}

func (u *UserController) addUser(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	name := req.FormValue("name")
	userResult := u.userService.AddUser(name)

	if userResult.Err != nil {
		http.Error(res, userResult.Err.Error(), userResult.StatusCode)
		return
	}

	if err := u.templates.ExecuteTemplate(res, "user", userResult.ResultData); err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}
}

func (u *UserController) getUsers(res http.ResponseWriter, req *http.Request) {

}
