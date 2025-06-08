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

func (r repository) InsertUser(ctx context.Context, user model.UserModel) (model.UserModel, error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	qres := trx.Create(&user).Error

	return user, qres
}

func (r repository) InsertUserDetail(ctx context.Context, ud model.UserDetailModel) error {
	trx := transaction.GetTrxContext(ctx, r.db)
	qres := trx.Create(&ud).Error

	return qres
}

func (r repository) InsertUserPreference(ctx context.Context, up model.UserPreferenceModel) error {
	trx := transaction.GetTrxContext(ctx, r.db)
	qres := trx.Create(&up).Error

	return qres
}

func (r repository) GetUserByUsername(ctx context.Context, username string) (user []model.UserModel, err error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	err = trx.Select("id, username, created_at, updated_at").Where("username = ?", username).Find(&user).Error
	return user, err
}

func (r repository) GetPasswordByUsername(ctx context.Context, username string) (user []model.UserModel, err error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	err = trx.Select("id, password, username, created_at, updated_at").Where("username = ?", username).Find(&user).Error
	return user, err
}

func (r repository) GetUserDetailById(ctx context.Context, id string) (user model.UserDetailModel, err error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	err = trx.Where("user_id = ?", id).First(&user).Error
	return user, err
}

func (r repository) GetUserPreference(ctx context.Context, userID string) (res model.UserPreferenceModel, err error) {
	trx := transaction.GetTrxContext(ctx, r.db)
	err = trx.Where("user_id = ?", userID).First(&res).Error
	return res, err
}

// func (r repository) PutUser(ctx context.Context, ud model.UserDetailModel, up model.UserPreferenceModel) (resultUD model.UserDetailModel, resultUP model.UserPreferenceModel, err error) {
// 	return ud, up, nil
// }

// func (r repository) GetUser(ctx context.Context, userID string) (resultUD model.UserDetailModel, resultUP model.UserPreferenceModel, err error) {
// 	var (
// 		ud model.UserDetailModel
// 		up model.UserPreferenceModel
// 	)
// 	return ud, up, nil
// }

// func (r repository) GetUserStatistic(ctx context.Context, userID string) (us model.UserStatisticModel, err error) {
// 	var user model.UserStatisticModel
// 	return user, nil
// }
