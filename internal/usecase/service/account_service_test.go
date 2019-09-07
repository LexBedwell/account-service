package service

import (
	"testing"
	"github.com/golang/mock/gomock"
	"github.com/lexbedwell/account-service/mocks"
	
)

func TestGetPongFromPing(t *testing.T) {
    mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	
	mockDaoInterface := mocks.NewMockDaoInterface(mockCtrl)
	
	var svc AccountService
	svc = AccountService{Dao: mockDaoInterface}
	
	serviceResponse := svc.GetPongFromPing()

	if(serviceResponse != "pong") {
		t.Errorf("Expected serviceResponse to be pong got %v", serviceResponse)
	}

}

func TestSvcGetUserFromId(t *testing.T) {
    mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	
	mockDaoInterface := mocks.NewMockDaoInterface(mockCtrl)

	mockDaoInterface.EXPECT().GetUserFromId("1").Return("lexbedwell@gmail.com", nil).Times(1)
	
	var svc AccountService
	svc = AccountService{Dao: mockDaoInterface}
	
	serviceResponse, err := svc.GetInfoFromId("1")

	if err != nil {
		t.Errorf("Error %v", err.Error())
	}

	if(serviceResponse != "lexbedwell@gmail.com") {
		t.Errorf("Expected serviceResponse to be lexbedwell@gmail.com got %v", serviceResponse)
	}

}

func TestSvcPostUser(t *testing.T) {
    mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	
	mockDaoInterface := mocks.NewMockDaoInterface(mockCtrl)

	mockDaoInterface.EXPECT().CreateUser("newUser@newEmail.com").Return("2", nil).Times(1)
	
	var svc AccountService
	svc = AccountService{Dao: mockDaoInterface}
	
	serviceResponseId, serviceResponseEmail, err := svc.PostUser("newUser@newEmail.com")

	if err != nil {
		t.Errorf("Error %v", err.Error())
	}

	if(serviceResponseId != "2") {
		t.Errorf("Expected serviceResponseId to be 2 got %v", serviceResponseId)
	}

	if(serviceResponseEmail != "newUser@newEmail.com") {
		t.Errorf("Expected serviceResponseEmail to be newUser@newEmail.com got %v", serviceResponseEmail)
	}

}

