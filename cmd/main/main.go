package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	calculatemaxprofitcontroller "maxprofit/internal/controller/calculatemaxprofit"
	calculatemaxprofitinteractor "maxprofit/internal/interactor/calculatemaxprofit"
	"maxprofit/internal/jwt"
	"maxprofit/internal/middleware"
)

const (
	maxProfitEndpoint = "/profit/max"
	jwtSecretEnvKey   = "JWT_SECRET"
	v1                = "/v1"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// TODO build rpc
	// TODO use structured logger [maybe slog]
	// TODO unit-tests

	// [1:] to ignore first arg which is preserved for exec call [e.g. './main.exe']
	args := os.Args[1:]
	fmt.Printf("args: %v", args)

	secretKey := os.Getenv(jwtSecretEnvKey)
	jwtV5Validator := jwt.NewV5Validator(secretKey)
	// TODO replace with DI [either to use Wire lib, or manual registry]
	calculateMaxProfitInteractor := calculatemaxprofitinteractor.NewInteractor()
	calculateMaxProfitController := calculatemaxprofitcontroller.NewController(calculateMaxProfitInteractor)

	r := gin.Default()
	r.Use(middleware.JWTValidatorHandlerFunc(jwtV5Validator))
	// TODO wrap the apis with docs [OpenAi-3.0 specs swagger, (either using swago or ogen)]
	v1Router := r.Group(v1)
	{
		v1Router.POST(maxProfitEndpoint, calculateMaxProfitController.MaxProfitHandler)
	}

	// listen and serve on 0.0.0.0:8080
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
