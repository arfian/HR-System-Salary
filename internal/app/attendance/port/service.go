package port

import (
	"context"
	"hr-system-salary/internal/app/attendance/model"
)

type IAttendanceService interface {
	AddAttendanceEmployee(ctx context.Context, username string) (res *model.AttendanceModel, err error)

	// BulkInserAttendance(ctx context.Context, attendances []model.AttendanceModel) error
}
