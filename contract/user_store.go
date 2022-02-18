package contract

import (
	"context"

	"github.com/SajjadManafi/simple-uber/models"
)

type UserStore interface {
	CreateUser(ctx context.Context, arg models.CreateUserParams) (models.User, error)
	GetUser(ctx context.Context, id int32) (models.User, error)
	ListUsers(ctx context.Context, arg models.ListUsersParams) ([]models.User, error)
	UpdateUser(ctx context.Context, arg models.UpdateUserParams) (models.User, error)
	DeleteUser(ctx context.Context, id int32) error
	AddUserBalance(ctx context.Context, arg models.AddUserBalanceParams) (models.User, error)
}
