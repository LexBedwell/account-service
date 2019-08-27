package adapters

import (
	"context"
	"net/http"
	"encoding/json"
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