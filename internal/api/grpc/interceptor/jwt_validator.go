package interceptor

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"maxprofit/internal/jwt"
)

const (
	authorizationMetadataKey = "x-authorization"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	// Unauthenticated is equivalent to 401 Unauthorized which means that the credentials are invalid
	// PermissionDenied is equivalent to 403 Forbidden which means that the credentials are valid however not permitted to perform such an action [e.g. role based]
	errInvalidToken = status.Errorf(codes.Unauthenticated, "invalid token")
)

func UnaryServerJwtValidatorFunc(jwtValidator jwt.Validator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errMissingMetadata
		}

		token := md.Get(authorizationMetadataKey)
		if len(token) == 0 {
			return nil, errInvalidToken
		}

		if isValid, err := jwtValidator.ValidateToken(token[0]); err != nil || !isValid {
			log.Printf("error malformed or invalid token %s", err.Error())
			return nil, errInvalidToken
		}

		log.Println("authorized")
		return handler(ctx, req)
	}
}
