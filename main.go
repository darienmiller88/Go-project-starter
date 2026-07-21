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
	fmt.Println("Hello world")

	app := chi.NewRouter()


}