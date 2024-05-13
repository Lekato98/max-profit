package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	calculatemaxprofitcontroller "maxprofit/internal/controller/calculatemaxprofit"
	calculatemaxprofitinteractor "maxprofit/internal/interactor/calculatemaxprofit"
	"maxprofit/internal/jwt"
	"maxprofit/internal/middleware"
	"maxprofit/internal/proto"
)

const (
	maxProfitEndpoint = "/profit/max"
	jwtSecretEnvKey   = "JWT_SECRET"
	v1                = "/v1"
)

type calculateMaxProfitServer struct {
	proto.UnimplementedProfitServer
	calculateMaxProfitInteractor *calculatemaxprofitinteractor.Interactor
}

func newCalculateMaxProfitServer(calculateMaxProfitInteractor *calculatemaxprofitinteractor.Interactor) *calculateMaxProfitServer {
	return &calculateMaxProfitServer{
		calculateMaxProfitInteractor: calculateMaxProfitInteractor,
	}
}

// TODO protect the endpoint with JWT validator check where we need to pass it
func (c *calculateMaxProfitServer) CalculateMaxProfit(ctx context.Context, profits *proto.Profits) (*proto.MaxProfit, error) {
	maxProfit, err := c.calculateMaxProfitInteractor.CalculateMaxProfit(ctx, profits.GetValues())
	if err != nil {
		return nil, err
	}

	return &proto.MaxProfit{
		Value: maxProfit,
	}, nil
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// TODO clean-up main and move RPC routes
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

	errChan := make(chan error, 2)
	go func() {
		// listen and serve on 0.0.0.0:8080
		errChan <- r.Run()
	}()

	go func() {
		// TODO use framework instead
		lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("RPC_PORT")))
		if err != nil {
			errChan <- err
			return
		}

		s := grpc.NewServer()
		proto.RegisterProfitServer(s, newCalculateMaxProfitServer(calculateMaxProfitInteractor))
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
