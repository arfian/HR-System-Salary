package port

import (
	"context"
	"hr-system-salary/internal/app/user/model"
	"hr-system-salary/internal/app/user/payload"
)

type IUserService interface {
	Register(ctx context.Context, user model.UserModel, ud model.UserDetailModel, up model.UserPreferenceModel) (token string, err error)

	Login(ctx context.Context, user model.UserModel) (token string, err error)

	GetUser(ctx context.Context, username string) (res *payload.User, err error)

	// GetUserStatistic(ctx context.Context, userID string) (us model.UserStatisticModel, err error)
}
