package auth

import (
	"context"
)

// ServicePort encapsulates the authentication logic.
type ServicePort interface {
	// // authenticate authenticates a user using phone number and pin.
	// // It returns a JWT token if authentication succeeds. Otherwise, an error is returned.
	Login(ctx context.Context, req RequestLogin) (ResponseLogin, error)
	// RefreshToken refresh the access token
	RefreshToken(ctx context.Context, req RequestRefreshToken) (ResponseLogin, error)
}

// Identity represents an authenticated user iddomain.
type Identity interface {
	// GetID returns the user ID.
	GetID() string
	// GetUsername returns the username.
	GetUsername() string
	// GetPassword returns password
	GetPassword() string
}
