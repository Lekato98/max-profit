package interceptor

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"maxprofit/internal/jwt"
)

const (
	authorizationMetadataKey = "x-authorization"
)

func UnaryServerJwtValidatorFunc(jwtValidator jwt.Validator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		defer func() {
			if r := recover(); r != nil {
				debug.PrintStack()
			}
		}()

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			err := fmt.Errorf("something went wrong while extracting metadata from incoming context")
			log.Printf(err.Error())
			return nil, status.Error(codes.InvalidArgument, err.Error())
		}

		token := md.Get(authorizationMetadataKey)[0]
		if isValid, err := jwtValidator.ValidateToken(token); err != nil || !isValid {
			err := fmt.Errorf("invalid token unauthorized %w", err)
			log.Printf(err.Error())
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}

		return handler(ctx, req)
	}
}
