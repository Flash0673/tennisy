package auth

import (
	"errors"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

var whiteList = map[string]bool{
	"/v1/login":    true,
	"/v1/register": true,
	"/v1/refresh":  true,
}

type Parser interface {
	Parse(token string) (string, error)
}

func NewAuthMiddleware(parser Parser) func(next runtime.HandlerFunc) runtime.HandlerFunc {
	return func(next runtime.HandlerFunc) runtime.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			if whiteList[r.URL.Path] {
				next(w, r, pathParams)
				return
			}

			// Example: Check if token exists
			token := r.Header.Get("Authorization")
			if token == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			uid, err := parseToken(token, parser)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			r.Header.Set("Grpc-Metadata-User-ID", uid)

			// Pass the request to the gRPC-Gateway mux
			next(w, r, pathParams)
		}
	}
}

func parseToken(token string, parser Parser) (string, error) {

	headerParts := strings.Split(token, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return "", errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return "", errors.New("token is empty")
	}

	return parser.Parse(headerParts[1])
}
