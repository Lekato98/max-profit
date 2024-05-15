package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"maxprofit/config"
	"maxprofit/internal/api/grpc/interceptor"
	"maxprofit/internal/api/grpc/service/profit"
	calculatemaxprofitcontroller "maxprofit/internal/api/http/controller/profit"
	"maxprofit/internal/api/http/middleware"
	calculatemaxprofitinteractor "maxprofit/internal/interactor/profit"
	"maxprofit/internal/jwt"
)

const (
	maxProfitEndpoint = "/profit/max"
	v1                = "/v1"
)

func main() {
	// TODO clean-up main and move RPC routes
	// TODO use structured logger [maybe slog]
	// TODO unit-tests

	// [1:] to ignore first arg which is preserved for exec call [e.g. './main.exe']
	args := os.Args[1:]
	fmt.Printf("args: %v", args)
	configs, err := config.LoadConfigs()
	if err != nil {
		log.Fatal(err)
	}

	jwtV5Validator := jwt.NewV5Validator(configs.SecretKey)
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

	errChan := make(chan error, 2)
	go func() {
		// listen and serve on 0.0.0.0:8080
		errChan <- r.Run()
	}()

	go func() {
		// TODO maybe use framework
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", configs.RpcPort))
		defer lis.Close()
		if err != nil {
			errChan <- err
			return
		}

		serverChain := grpc.ChainUnaryInterceptor(interceptor.UnaryServerJwtValidatorFunc(jwtV5Validator))
		s := grpc.NewServer(serverChain)
		profitServer := profit.NewProfitServer(calculateMaxProfitInteractor)
		profit.RegisterProfitServer(s, profitServer)
		log.Printf("Listening and serving RPC on %s", lis.Addr())
		errChan <- s.Serve(lis)
	}()

	for err := range errChan {
		log.Printf("reading from errChan")
		if err != nil {
			log.Fatal(err)
		}
	}
}
