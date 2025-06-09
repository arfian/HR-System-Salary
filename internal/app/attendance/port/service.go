package port

import (
	"context"
	"hr-system-salary/internal/app/attendance/model"
	"hr-system-salary/internal/app/attendance/payload"
)

type IAttendanceService interface {
	AddAttendanceEmployee(ctx context.Context, username string) (res *model.AttendanceModel, err error)

	BulkInserAttendance(ctx context.Context, param payload.ParamBulkAttendance, username string) error

	AddOvertime(ctx context.Context, param payload.ParamOvertime, username string) (*model.OvertimeModel, error)
}
