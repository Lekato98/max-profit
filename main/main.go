package main

import (
	"awesomeProject1/internal/controller"
	"awesomeProject1/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	route = "v1/profit/max"
)

func main() {
	r := gin.Default()
	if err := godotenv.Load("./config/.env"); err != nil {
		log.Fatal(err)
	}
	secretKey := os.Getenv("JWT_SECRET")
	r.Use(middleware.JWTValidatorFuncHandler(secretKey))
	r.POST(route, controller.CalculateMaxProfitHandler)
	err := r.Run() // listen and serve on 0.0.0.0:8080
	log.Fatal(err)
}
