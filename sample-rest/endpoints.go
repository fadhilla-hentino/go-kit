package main

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// UserRequest define request struct
type UserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// UserResponse define response struct
type UserResponse struct {
	UserID string `json:"userID"`
	Error  error  `json:"error"`
}

// UserInfoRequest define request struct
type UserInfoRequest struct {
	UserID string `json:"userID"`
}

// UserInfoResponse define response for Get User Info struct
type UserInfoResponse struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Error    error  `json:"error"`
}

type Endpoints struct {
	UserEndpoint    endpoint.Endpoint
	GetUserEndpoint endpoint.Endpoint
}

// NewEndpoint make endpoint
func NewEndpoint(svc Service) *Endpoints {
	return &Endpoints{
		UserEndpoint:    NewUserEndpoint(svc),
		GetUserEndpoint: NewGetUserEndpoint(svc),
	}
}

// NewUserEndpoint make endpoint
func NewUserEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UserRequest)

		userID, err := svc.Register(ctx, &req)

		return UserResponse{
			UserID: userID,
			Error:  err,
		}, nil
	}
}

// NewGetUserEndpoint make endpoint
func NewGetUserEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*UserInfoRequest)

		user, err := svc.GetUserByID(ctx, req.UserID)
		if err != nil {
			user.Error = err
		}
		return user, nil
	}
}
