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
	rpc               = "rpc"
)

func main() {
	// TODO clean-up main and move RPC routes
	// TODO use structured logger [maybe slog]
	// TODO unit-tests
	configs, err := config.LoadConfigs()
	if err != nil {
		log.Fatal(err)
	}

	// [1:] to ignore first arg which is preserved for exec call [e.g. './main.exe']
	args := os.Args[1:]
	fmt.Printf("prased args: %v\n", args)
	serverType := args[0]
	jwtV5Validator := jwt.NewV5Validator(configs.SecretKey)
	// TODO replace with DI [either to use Wire lib, or manual registry]
	calculateMaxProfitInteractor := calculatemaxprofitinteractor.NewInteractor()

	switch serverType {
	case rpc:
		// TODO maybe use framework
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", configs.RpcPort))
		if err != nil {
			log.Fatal(err)
		}

		defer lis.Close()
		serverChain := grpc.ChainUnaryInterceptor(
			interceptor.UnaryServerRecovery,
			interceptor.UnaryServerJwtValidatorFunc(jwtV5Validator),
		)
		s := grpc.NewServer(serverChain)

		profitServer := profit.NewProfitServer(calculateMaxProfitInteractor)
		profit.RegisterProfitServer(s, profitServer)

		log.Printf("Listening and serving RPC on %s", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatal(err)
		}
	default:
		calculateMaxProfitController := calculatemaxprofitcontroller.NewController(calculateMaxProfitInteractor)

		r := gin.Default()
		r.Use(middleware.JWTValidatorHandlerFunc(jwtV5Validator))
		// TODO wrap the apis with docs [OpenAi-3.0 specs swagger, (either using swago or ogen)]
		v1Router := r.Group(v1)
		{
			v1Router.POST(maxProfitEndpoint, calculateMaxProfitController.MaxProfitHandler)
		}

		if err := r.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
