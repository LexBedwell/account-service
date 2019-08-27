package handlers

import (
	"context"
	encoderAdapters "github.com/lexbedwell/account-service/internal/usecase/models/responses"
	decoderAdapters "github.com/lexbedwell/account-service/internal/usecase/models/requests"
	"github.com/lexbedwell/account-service/internal/adapters/transport/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/lexbedwell/account-service/internal/usecase/service"
)

var svc service.AccountService

type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)

func NewGetPongFromPingHandler(svc service.AccountService) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.MakeGetPongFromPingEndpoint(svc),
		decoderAdapters.DecodeGetPongFromPingRequest,
		encoderAdapters.EncodeResponse,
	)
}

func NewPostUserHandler(svc service.AccountService) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.MakePostUserEndpoint(svc),
		decoderAdapters.DecodePostUserRequest,
		encoderAdapters.EncodeResponse,
	)
}

func NewGetInfoFromIdHandler(svc service.AccountService) *httptransport.Server {
	return httptransport.NewServer(
		endpoints.MakeGetInfoFromIdEndpoint(svc),
		decoderAdapters.DecodeGetInfoFromIdRequest,
		encoderAdapters.EncodeResponse,
	)
}