package service

import (
	"context"
	"errors"
	"time"

	"hr-system-salary/internal/app/reimbursement/model"
	"hr-system-salary/internal/app/reimbursement/payload"
	"hr-system-salary/internal/app/reimbursement/port"
	userPort "hr-system-salary/internal/app/user/port"
)

type service struct {
	reimbursementRepo port.IReimbursementRepository
	userRepo          userPort.IUserRepository
}

func New(reimbursementRepo port.IReimbursementRepository, userRepo userPort.IUserRepository) port.IReimbursementService {
	return &service{
		reimbursementRepo: reimbursementRepo,
		userRepo:          userRepo,
	}
}

func (s *service) AddReimbursement(ctx context.Context, username string, param payload.ParamReimbursement) (res *model.ReimbursementModel, err error) {
	users, qerr := s.userRepo.GetUserByUsername(ctx, username)
	if len(users) == 0 || qerr != nil {
		return nil, errors.New("user not found")
	}

	reimbursementDate, _ := time.Parse("2006-01-02", param.ReimbursementDate)
	reimbursement := model.ReimbursementModel{
		Employee:            users[0].ID.String(),
		ReimbursementDate:   reimbursementDate,
		Status:              "NOT PAID",
		CreatedBy:           username,
		Description:         param.Description,
		ReimbursementType:   param.ReimbursementType,
		ReimbursementAmount: param.ReimbursementAmount,
	}
	reimbursement, qerr = s.reimbursementRepo.InsertReimbursement(ctx, reimbursement)
	if qerr != nil {
		return nil, qerr
	}

	return &reimbursement, nil
}
