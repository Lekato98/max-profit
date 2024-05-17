package interceptor

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnaryServerRecovery(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	defer func() {
		if r := recover(); r != nil {
			stackTrace := debug.Stack()
			log.Println(fmt.Errorf("panic occurred: %v\n%s", r, stackTrace))

			resp = nil
			err = status.Error(codes.Unknown, "internal server error")
		}
	}()

	return handler(ctx, req)
}
