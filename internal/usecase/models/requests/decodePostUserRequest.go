package adapters

import (
	"context"
	"encoding/json"
	"net/http"
)

type PostUserRequest struct {
	Email string `json:"email"`
}

func DecodePostUserRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request PostUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}