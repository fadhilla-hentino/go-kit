package main

import (
	"context"
	"encoding/json"
	"strings"

	"fadhilla-hentino/go-kit/sample-rest/repository"
	"github.com/google/uuid"
)

// Service Define a service interface
type Service interface {

	// Register register user
	Register(ctx context.Context, userRequest *UserRequest) (string, error)

	// GetUserByID get user by user ID
	GetUserByID(ctx context.Context, id string) (*UserInfoResponse, error)
}

//UserService implement Service interface
type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// Register implement Register method
func (s UserService) Register(ctx context.Context, userRequest *UserRequest) (string, error) {

	userID := generateUUID()
	bytes, err := json.Marshal(&userRequest)
	if err != nil {
		return "", err
	}

	return userID, s.userRepo.Store(ctx, userID, string(bytes))
}

// GetUserByID implement GetUserByID method
func (s UserService) GetUserByID(ctx context.Context, id string) (*UserInfoResponse, error) {

	rawUser, err := s.userRepo.Load(ctx, id)
	if err != nil {
		return nil, err
	}

	var resp *UserInfoResponse
	if err = json.Unmarshal([]byte(rawUser), &resp); err != nil {
		return nil, err
	}
	resp.UserID = id

	return resp, nil
}

// generateUUID generates UUID v4 without hyphens
func generateUUID() string {
	id, _ := uuid.NewRandom()
	return strings.ReplaceAll(id.String(), "-", "")
}
