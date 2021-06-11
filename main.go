package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/pmqueiroz/http-advices/routers"
)

func main() {
	godotenv.Load()

	port := os.Getenv("PORT")

	routers.Run(&port)
}