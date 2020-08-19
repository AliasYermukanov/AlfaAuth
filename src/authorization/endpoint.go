package authorization

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func createUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createUserRequest)
		resp, err := s.createUser(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type createUserRequest struct {
	MobilePhone string `json:"mobile_phone"`
	Password    string `json:"password"`
}

type createUserResponse struct {
	Message string `json:"message"`
}

func authenticationEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(createUserRequest)
		resp, err := s.authenticateUser(&req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}

type authenticateUserResponse struct {
	Uid string `json:"uid"`
}

func UpdateUserEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		s.updateUser()
		return nil, nil
	}
}
