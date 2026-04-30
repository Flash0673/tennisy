package metadata

import (
	"context"
	"net/http"

	"google.golang.org/grpc/metadata"
)

func UserMetadata(ctx context.Context, r *http.Request) metadata.MD {
	// Читаем заголовок, который мы установили в Middleware
	uID := r.Header.Get("Grpc-Metadata-User-ID")
	return metadata.Pairs("user_id", uID)
}
