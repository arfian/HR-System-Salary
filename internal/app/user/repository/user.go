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
