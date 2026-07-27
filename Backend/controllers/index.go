package controllers

import (
	"GoProjectStarter/Backend/repositories"
	"GoProjectStarter/Backend/services"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type IndexController struct{
	Router *chi.Mux
	db     *sqlx.DB
}

func NewIndexController(db *sqlx.DB) *IndexController {
	c := &IndexController{
		Router: chi.NewRouter(),
		db:     db,
	}

	c.registerRoutes()

	return c
}

// registerRoutes sets up the routes for the IndexController by mounting all sub-routes in controllers.
func (c *IndexController) registerRoutes() {
	vc := NewViewsController()
	uc := NewUserController(services.NewUserService(repositories.NewUserRepository(c.db)))

	c.Router.Mount("/", vc.Router)
	c.Router.Mount("/users", uc.Router)
}