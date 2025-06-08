package port

import (
	"context"
	"hr-system-salary/internal/app/user/model"
)

type IUserRepository interface {
	InsertUser(ctx context.Context, user model.UserModel) (model.UserModel, error)

	InsertUserDetail(ctx context.Context, ud model.UserDetailModel) error

	InsertUserPreference(ctx context.Context, up model.UserPreferenceModel) error

	GetUserByUsername(ctx context.Context, username string) (user []model.UserModel, err error)

	GetPasswordByUsername(ctx context.Context, username string) (user []model.UserModel, err error)

	GetUserDetailById(ctx context.Context, id string) (user model.UserDetailModel, err error)

	GetUserPreference(ctx context.Context, userID string) (res model.UserPreferenceModel, err error)

	// GetUserStatistic(ctx context.Context, userID string) (us model.UserStatisticModel, err error)
}
