package interceptors

import (
	"context"
	"golara/app/models"
	"golara/core/facades/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

type AuthInterceptor struct{}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, nil
		}

		authVal := md["authorization"]
		if len(authVal) == 0 {
			return nil, nil
		}
		token := authVal[0]
		splitToken := strings.Split(token, "Bearer ")
		if len(splitToken) != 2 {
			return nil, nil
		}
		token = splitToken[1]
		if token == "" {
			return nil, nil
		}

		var user models.User
		err := user.GetUserByToken(token)
		if err != nil {
			return nil, nil
		}

		auth.LoggedInUser(user)

		return handler(ctx, req)
	}
}
