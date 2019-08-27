package adapters

import (
	"context"
	"net/http"
)

type GetInfoFromIdRequest struct {
	Id string 
}

func DecodeGetInfoFromIdRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request GetInfoFromIdRequest
	query := r.URL.Query()
	request = GetInfoFromIdRequest{query.Get("id")}
	return request, nil
}