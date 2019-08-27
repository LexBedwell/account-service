package adapters

import (
	"context"
	"net/http"
)

func DecodeGetPongFromPingRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	type getPingRequest struct {}
	var request getPingRequest
	return request, nil
}