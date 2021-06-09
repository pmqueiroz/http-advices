package main

import (
	"fmt"
	"os"

	"github.com/pmqueiroz/http-advices/routers"
)

func main() {
	port := os.Getenv("PORT")

	fmt.Println(port)
	routers.Run()
}