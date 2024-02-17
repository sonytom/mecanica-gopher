package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sonytom/mecanica-gopher/src/controller/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println(os.Getenv("TESTE"))

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	err = router.Run(os.Getenv("PORT"))
	if err != nil {
		log.Fatalln(err)
	}

}
