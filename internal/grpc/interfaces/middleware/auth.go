package middleware

import (
	"context"
	"fmt"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
	"net/http"
)

type key int

var result key

// AuthFunc is ...
func AuthFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "Authorization")
	if err != nil {
		return nil, err
	}
	fmt.Printf("receive token: %s\n", token)
	err = validateToken(token)
	if err != nil {
		return nil, grpc.Errorf(http.StatusUnauthorized, "invalid token")
	}
	newCtx := context.WithValue(ctx, result, "ok")
	return newCtx, nil
}

// Check your authorization token is valid to ask authorization server
func validateToken(token string) error {
	return nil
}
