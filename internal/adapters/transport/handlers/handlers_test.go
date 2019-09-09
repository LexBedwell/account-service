package handlers

import (
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/lexbedwell/account-service/mocks"
	"github.com/lexbedwell/account-service/internal/usecase/service"
	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http/httptest"	
	"net/http"
	"strings"
)

func TestGetPongFromPingRoute(t *testing.T) {
    mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	
	mockDaoInterface := mocks.NewMockDaoInterface(mockCtrl)
	
	var svc service.AccountService
	svc = service.AccountService{Dao: mockDaoInterface}

	handler := NewGetPongFromPingHandler(&svc)

	responseRecorder := httptest.NewRecorder()
	
	request := httptest.NewRequest("GET", "/ping", nil)
	handler.ServeHTTP(responseRecorder, request)

	resp := responseRecorder.Result()
	
	if status := resp.StatusCode; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
	}
	
    if strings.TrimSpace(responseRecorder.Body.String()) != `{"response":"pong"}` {
        t.Errorf("handler returned unexpected body: got %v want %v",
		strings.TrimSpace(responseRecorder.Body.String()), `{"response":"pong"}`)
    }

}

func TestSvcGetUserFromIdRoute(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	
	mockDaoInterface := mocks.NewMockDaoInterface(mockCtrl)

	mockDaoInterface.EXPECT().GetUserFromId("1").Return("lexbedwell@gmail.com", nil).Times(1)
	
	var svc service.AccountService
	svc = service.AccountService{Dao: mockDaoInterface}

	handler := NewGetInfoFromIdHandler(&svc)

	responseRecorder := httptest.NewRecorder()
	
	request := httptest.NewRequest("GET", "/user/?id=1", nil)
	handler.ServeHTTP(responseRecorder, request)

	resp := responseRecorder.Result()
	
	if status := resp.StatusCode; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
	}
	
    if strings.TrimSpace(responseRecorder.Body.String()) != `{"id":"1","email":"lexbedwell@gmail.com","error":""}` {
        t.Errorf("handler returned unexpected body: got %v want %v",
		strings.TrimSpace(responseRecorder.Body.String()), `{"id":"1","email":"lexbedwell@gmail.com","error":""}`)
	}
}

func TestSvcPostUserRoute(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	
	mockDaoInterface := mocks.NewMockDaoInterface(mockCtrl)

	mockDaoInterface.EXPECT().CreateUser("newUser@newEmail.com").Return("2", nil).Times(1)
	
	var svc service.AccountService
	svc = service.AccountService{Dao: mockDaoInterface}

	handler := NewPostUserHandler(&svc)

	responseRecorder := httptest.NewRecorder()

	type PostUserRequest struct {
		Email string `json:"email"`
	}
	
	request := httptest.NewRequest("POST", "/create", strings.NewReader(`{"email":"newUser@newEmail.com"}`))
	handler.ServeHTTP(responseRecorder, request)

	resp := responseRecorder.Result()
	
	if status := resp.StatusCode; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
	}
	
    if strings.TrimSpace(responseRecorder.Body.String()) != `{"id":"2","email":"newUser@newEmail.com","error":""}` {
        t.Errorf("handler returned unexpected body: got %v want %v",
		strings.TrimSpace(responseRecorder.Body.String()), `{"id":"2","email":"newUser@newEmail.com","error":""}`)
	}

}

func TestGetMetricsRoute(t *testing.T) {

	handler := NewPrometheusHandler(promhttp.Handler())

	responseRecorder := httptest.NewRecorder()
	
	request := httptest.NewRequest("GET", "/metrics", nil)
	handler.ServeHTTP(responseRecorder, request)

	resp := responseRecorder.Result()
	
	if status := resp.StatusCode; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
	}

}