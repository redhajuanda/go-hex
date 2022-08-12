package mariadb

import (
	"context"
	"database/sql"
	"go-hex/internal/domain"
	"go-hex/pkg/otel"
	"go-hex/shared/ierr"

	"github.com/pkg/errors"
	"github.com/uptrace/bun"
)

// UserRepository encapsulates the logic to access users from the data source.
type UserRepository struct {
	db DBI
}

// NewUserRepository creates a new user repository
func NewUserRepository(db DBI) *UserRepository {
	return &UserRepository{db}
}

// GetByID returns the user with the specified user ID.
func (r *UserRepository) GetByID(ctx context.Context, userID string) (domain.User, error) {

	ctx, span := otel.Start(ctx)
	defer span.End()

	user := domain.User{ID: userID}
	err := r.db.
		NewSelect().
		Model(&user).
		Where("?=?", bun.Ident("id"), userID).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, ierr.ErrResourceNotFound
		}
		return domain.User{}, errors.Wrap(err, "cannot get user")
	}

	return user, nil
}

// GetByUsername returns the user with the specified username.
func (r *UserRepository) GetByUsername(ctx context.Context, username string) (domain.User, error) {

	ctx, span := otel.Start(ctx)
	defer span.End()

	var user domain.User

	err := r.db.
		NewSelect().
		Model(&user).
		Where("?=?", bun.Ident("username"), username).
		Scan(ctx)

	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{}, ierr.ErrResourceNotFound
		}
		return domain.User{}, errors.Wrap(err, "cannot get user")
	}

	return user, nil
}

// IsUserExistByID checks wether user exists
func (r *UserRepository) IsUserExistByID(ctx context.Context, userID string) (bool, error) {

	ctx, span := otel.Start(ctx)
	defer span.End()

	user := domain.User{ID: userID}
	exist, err := r.db.
		NewSelect().
		Model(&user).
		Where("?=?", bun.Ident("id"), userID).
		Exists(ctx)

	if err != nil {
		return false, errors.Wrap(err, "cannot check user")
	}
	return exist, nil
}

// IsUserExistByUsername checks whether user exists by username
func (r *UserRepository) IsUserExistByUsername(ctx context.Context, username string) (bool, error) {

	ctx, span := otel.Start(ctx)
	defer span.End()

	user := domain.User{}
	exist, err := r.db.
		NewSelect().
		Model(&user).
		Where("?=?", bun.Ident("username"), username).
		Exists(ctx)

	if err != nil {
		return false, errors.Wrap(err, "cannot check user")
	}
	return exist, nil
}

// Update updates the user with given ID in the storage.
func (r *UserRepository) Update(ctx context.Context, userID string, user domain.User) error {

	ctx, span := otel.Start(ctx)
	defer span.End()

	_, err := r.db.NewUpdate().
		Model(&user).
		OmitZero().
		Where("?=?", bun.Ident("id"), userID).
		Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "cannot update user")
	}
	return nil
}
