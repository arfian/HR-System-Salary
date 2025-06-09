package port

import (
	"context"
	"hr-system-salary/internal/app/attendance/model"
)

type IAttendanceRepository interface {
	InsertAttendanceEmployee(ctx context.Context, attendance model.AttendanceModel) (model.AttendanceModel, error)

	GetAttendanceByUserDate(ctx context.Context, userId string, attendanceDate string) (res []model.AttendanceModel, err error)

	GetDateRangeAttendanceByUser(ctx context.Context, userId string, startDate string, endDate string) (res []string, err error)

	BulkInsertAttendance(ctx context.Context, attendances []model.AttendanceModel) error

	GetSumOvertimeByUserDate(ctx context.Context, userId string, overtimeDate string) (res int, err error)

	InsertOvertime(ctx context.Context, overtime model.OvertimeModel) (model.OvertimeModel, error)

	GetAttendanceStatusByUserDate(ctx context.Context, userId string, attendanceDate string, status string) (res []model.AttendanceModel, err error)
}
