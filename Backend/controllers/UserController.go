package controllers

import (
	"GoProjectStarter/Backend/services"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type UserController struct {
	Router     *chi.Mux
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	
	uc := &UserController{
		Router:      chi.NewRouter(),
		userService: userService,
	}

	uc.registerViewRoutes()

	return uc
}

func (u *UserController) registerViewRoutes() {
	u.Router.Get("/add-user", u.addUser)
}

func (u *UserController) addUser(response http.ResponseWriter, request *http.Request) {
	
}