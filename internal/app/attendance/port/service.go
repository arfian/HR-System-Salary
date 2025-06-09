package port

import (
	"context"
)

type IAttendanceService interface {
	AddAttendanceEmployee(ctx context.Context, username string) error

	// BulkInserAttendance(ctx context.Context, attendances []model.AttendanceModel) error
}
