package user

import (
	"context"
	"go-hex/configs"
	"go-hex/internal/domain"
	"go-hex/internal/repository/port"
	"go-hex/pkg/auth"
	"go-hex/pkg/otel"
)

// Service encapsulates the user logic.
type Service struct {
	cfg         *configs.Config
	repoRegitry port.RepositoryRegistry
}

// NewService creates and returns a new user service
func NewService(cfg *configs.Config, repoRegitry port.RepositoryRegistry) Service {
	return Service{cfg, repoRegitry}
}

// Get returns the user with the specified user ID or username.
func (s Service) Get(ctx context.Context) (domain.User, error) {

	ctx, span := otel.Start(ctx)
	defer span.End()

	user := auth.GetLoggedInUser(ctx)

	repoUser := s.repoRegitry.GetUserRepository()
	return repoUser.GetByID(ctx, user.ID)
}
