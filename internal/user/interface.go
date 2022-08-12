package user

import (
	"context"
	"go-hex/internal/domain"
)

// ServicePort encapsulates usecase logic for users.
type ServicePort interface {
	// Get returns the user with the specified user ID or username.
	Get(ctx context.Context) (domain.User, error)
}
