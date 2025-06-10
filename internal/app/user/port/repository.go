package port

import (
	"context"
	"hr-system-salary/internal/app/user/model"
)

type IUserRepository interface {
	InsertUser(ctx context.Context, user model.AuthUserModel) (model.AuthUserModel, error)

	InsertEmployee(ctx context.Context, ud model.EmployeeModel) error

	GetUserByUsername(ctx context.Context, username string) (user []model.AuthUserModel, err error)

	GetPasswordByUsername(ctx context.Context, username string) (user []model.AuthUserModel, err error)

	GetEmployeeByUsername(ctx context.Context, username string) (user model.EmployeeModel, err error)

	UpdateLastLogin(ctx context.Context, user model.AuthUserModel) error

	GetAttendanceOvertimeByEmployee(ctx context.Context, limit int, pageNo int, year int, month int) (res []model.AttendanceOvertimeModel, err error)

	GetCountEmployee(ctx context.Context) (res int64, err error)
}
