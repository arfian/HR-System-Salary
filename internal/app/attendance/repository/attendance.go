package repository

import (
	"context"

	"hr-system-salary/config/db"
	"hr-system-salary/pkg/cache"
	"hr-system-salary/pkg/transaction"

	"hr-system-salary/internal/app/attendance/model"
	"hr-system-salary/internal/app/attendance/port"
)

type repository struct {
	db    *db.GormDB
	cache cache.ICache
}

func NewRepository(db *db.GormDB) port.IAttendanceRepository {
	return repository{db: db}
}

func (r repository) InsertAttendanceEmployee(ctx context.Context, attendance model.AttendanceModel) (model.AttendanceModel, error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	qres := trx.Save(&attendance).Error

	return attendance, qres
}

func (r repository) GetAttendanceByUserDate(ctx context.Context, userId string, attendanceDate string) (res []model.AttendanceModel, err error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	err = trx.Where("employee = ?", userId).Where("check_in::date = ?", attendanceDate).Find(&res).Error
	return res, err
}

func (r repository) GetDateRangeAttendanceByUser(ctx context.Context, userId string, startDate string, endDate string) (res []string, err error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	err = trx.Table("attendance").Select("TO_CHAR(check_in :: DATE, 'yyyy-mm-dd') AS date_check_in").Where("employee = ?", userId).Where("check_in::date >= ?", startDate).Where("check_out::date <= ?", endDate).Pluck("date_check_in", &res).Error
	return res, err
}

func (r repository) BulkInsertAttendance(ctx context.Context, attendances []model.AttendanceModel) error {
	trx := transaction.GetTrxContext(ctx, r.db)
	qres := trx.Create(&attendances).Error

	return qres
}
