package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/log"
	gokithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// MakeHttpHandler make http handler use mux
func MakeHttpHandler(ctx context.Context, userSvc *UserService, logger log.Logger) http.Handler {
	r := mux.NewRouter()

	options := []gokithttp.ServerOption{
		gokithttp.ServerErrorLogger(logger),
		gokithttp.ServerErrorEncoder(gokithttp.DefaultErrorEncoder),
	}

	e := NewEndpoint(userSvc)

	r.Methods("POST").Path("/users").Handler(gokithttp.NewServer(
		e.UserEndpoint,
		decodeRequest,
		encodeRequest,
		options...,
	))

	r.Methods("GET").Path("/users/{userID}").Handler(gokithttp.NewServer(
		e.GetUserEndpoint,
		decodeGetUserRequest,
		encodeGetUserRequest,
		options...,
	))

	return r
}

// decodeRequest decode request params to struct
func decodeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req UserRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// encodeRequest encode response to return
func encodeRequest(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

// decodeGetUserRequest decode request params to struct
func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	userID, ok := vars["userID"]
	if !ok {
		return nil, errors.New("invalid path parameter")
	}

	return &UserInfoRequest{
		UserID: userID,
	}, nil
}

// encodeGetUserRequest encode response to return
func encodeGetUserRequest(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
