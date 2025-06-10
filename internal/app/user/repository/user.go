package repository

import (
	"context"

	"hr-system-salary/config/db"
	"hr-system-salary/pkg/cache"
	"hr-system-salary/pkg/transaction"

	"hr-system-salary/internal/app/user/model"
	"hr-system-salary/internal/app/user/port"
)

type repository struct {
	db    *db.GormDB
	cache cache.ICache
}

func NewRepository(db *db.GormDB) port.IUserRepository {
	return repository{db: db}
}

func (r repository) InsertUser(ctx context.Context, user model.AuthUserModel) (model.AuthUserModel, error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	qres := trx.Create(&user).Error

	return user, qres
}

func (r repository) InsertEmployee(ctx context.Context, ud model.EmployeeModel) error {
	trx := transaction.GetTrxContext(ctx, r.db)
	qres := trx.Create(&ud).Error

	return qres
}

func (r repository) GetUserByUsername(ctx context.Context, username string) (user []model.AuthUserModel, err error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	err = trx.Select("id, username, created_at, updated_at").Where("username = ?", username).Find(&user).Error
	return user, err
}

func (r repository) GetPasswordByUsername(ctx context.Context, username string) (user []model.AuthUserModel, err error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	err = trx.Select("id, password, username, created_at, updated_at").Where("username = ?", username).Find(&user).Error
	return user, err
}

func (r repository) GetEmployeeByUsername(ctx context.Context, username string) (user model.EmployeeModel, err error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	err = trx.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r repository) UpdateLastLogin(ctx context.Context, user model.AuthUserModel) error {
	trx := transaction.GetTrxContext(ctx, r.db)
	err := trx.Model(&model.AuthUserModel{}).Where("username = ?", user.Username).Update("last_login", user.LastLogin).Error
	return err
}

func (r repository) GetAttendanceOvertimeByEmployee(ctx context.Context, limit int, pageNo int, year int, month int) (res []model.AttendanceOvertimeModel, err error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	err = trx.Raw("SELECT em.ID, em.username, em.fullname, em.salary_amount AS basic_salary, COUNT(a.employee) AS total_attendance, SUM(o.overtime_hours) AS sum_overtime, SUM(r.reimbursement_amount) as total_reimbursement FROM auth_user au INNER JOIN employee em ON au.username=em.username LEFT OUTER JOIN attendance a ON au.id=a.employee LEFT OUTER JOIN overtime o ON au.ID=o.employee LEFT OUTER JOIN reimbursement r ON r.employee=au.id WHERE a is NULL OR r is NULL OR o is NULL OR (EXTRACT(MONTH FROM r.reimbursement_date) = ? AND EXTRACT(YEAR FROM r.reimbursement_date) = ?) OR (EXTRACT(MONTH FROM a.check_in) = ? AND EXTRACT(YEAR FROM a.check_in) = ?) OR (EXTRACT(MONTH FROM o.overtime_date) = ? AND EXTRACT(YEAR FROM o.overtime_date) = ?) GROUP BY em.ID OFFSET ? LIMIT ?", month, year, month, year, month, year, pageNo, limit).Scan(&res).Error
	return res, err
}

func (r repository) GetAllEmployee(ctx context.Context, limit int, pageNo int) (res []model.AttendanceOvertimeModel, err error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	err = trx.Raw("SELECT em.ID, em.username, em.fullname, em.salary_amount AS basic_salary FROM auth_user au INNER JOIN employee em ON au.username=em.username GROUP BY em.ID OFFSET ? LIMIT ?", pageNo, limit).Scan(&res).Error
	return res, err
}

func (r repository) GetCountEmployee(ctx context.Context) (res int64, err error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	var count int64 = 0
	err = trx.Model(&model.EmployeeModel{}).Count(&count).Error
	return count, err
}
