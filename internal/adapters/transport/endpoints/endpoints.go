package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/lexbedwell/account-service/internal/usecase/service"
	"github.com/lexbedwell/account-service/internal/usecase/models/requests"
)

func MakeGetInfoFromIdEndpoint(svc service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		type getIdResponse struct {
			Id   string `json:"id"`
			Email   string `json:"email"`
			Err string `json:"error"`
		}
		req := request.(adapters.GetInfoFromIdRequest)
		v, _ := svc.GetInfoFromId(req.Id)
		return getIdResponse{req.Id, v, ""}, nil
	}
}

func MakeGetPongFromPingEndpoint(svc service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		type getPingResponse struct {
			Response   string `json:"response"`
		}
		v := svc.GetPongFromPing()
		return getPingResponse{Response: v}, nil
	}
}

func MakePostUserEndpoint(svc service.AccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		type postUserResponse struct {
			Id   string `json:"id"`
			Email   string `json:"email"`
			Err string `json:"error"`
		}
		req := request.(adapters.PostUserRequest)
		id, email, err := svc.PostUser(req.Email)
		if err != nil {
			return postUserResponse{Id: "", Email: "", Err: err.Error()}, nil
		}
		return postUserResponse{Id: id, Email: email, Err: ""}, nil
	}
}