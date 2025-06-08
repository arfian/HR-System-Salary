package port

import (
	"context"
	"hr-system-salary/internal/app/user/model"
	"hr-system-salary/internal/app/user/payload"
)

type IUserService interface {
	Register(ctx context.Context, user model.AuthUserModel, em model.EmployeeModel) (token string, err error)

	Login(ctx context.Context, user model.AuthUserModel) (token string, err error)

	GetUser(ctx context.Context, username string) (res *payload.User, err error)
}
