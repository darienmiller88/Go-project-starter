package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"GoProjectStarter/Backend/database"
	"GoProjectStarter/Backend/controllers"
)

func main(){
	godotenv.Load()
	database.ConnectToSQL()

	router := chi.NewRouter()
	indexController := controllers.NewIndexController()

	router.Mount("/", indexController.Router)
	
	//Serve static files along the "/static" route
	fs := http.FileServer(http.Dir("static"))
	router.Handle("/static/*", http.StripPrefix("/static/", fs))

	port := os.Getenv("PORT")

	fmt.Println("Server is running on port:", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}