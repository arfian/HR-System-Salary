package service

import (
	"context"
	"errors"
	"time"

	"hr-system-salary/internal/app/user/model"
	"hr-system-salary/internal/app/user/payload"
	"hr-system-salary/internal/app/user/port"

	"hr-system-salary/config"
	"hr-system-salary/pkg/encrypt"

	jwt "github.com/golang-jwt/jwt/v5"
)

type service struct {
	userRepo port.IUserRepository
}

func New(userRepo port.IUserRepository) port.IUserService {
	return &service{
		userRepo: userRepo,
	}
}

func (s *service) Register(ctx context.Context, user model.AuthUserModel, em model.EmployeeModel) (token string, err error) {
	username, qerr := s.userRepo.GetUserByUsername(ctx, user.Username)
	if qerr != nil {
		return "", qerr
	}
	if len(username) > 0 {
		return "", errors.New("user already exists")
	}

	hash, qerr := encrypt.HashPassword(user.Password)
	if qerr != nil {
		return "", qerr
	}

	em.Username = user.Username
	em.CreatedBy = user.Username
	user.CreatedBy = user.Username
	user.LastLogin = time.Now()
	user.Password = hash
	user, qerr = s.userRepo.InsertUser(ctx, user)
	if qerr != nil {
		return "", qerr
	}

	qerr = s.userRepo.InsertEmployee(ctx, em)
	if qerr != nil {
		return "", qerr
	}

	tokenString, err := createToken(user, em)

	return tokenString, err
}

func createToken(user model.AuthUserModel, em model.EmployeeModel) (string, error) {
	configData := config.GetConfig()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"fullname": em.Fullname,
		"rolename": em.Rolename,
		"exp":      time.Now().Add(time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	})
	tokenString, err := claims.SignedString([]byte(configData.JWT.SigningKey))

	return tokenString, err
}

func (s service) Login(ctx context.Context, user model.AuthUserModel) (token string, err error) {
	users, qerr := s.userRepo.GetPasswordByUsername(ctx, user.Username)
	if len(users) == 0 || qerr != nil {
		return "", errors.New("incorrect username or password")
	}

	match := encrypt.CheckPasswordHash(user.Password, users[0].Password)
	if !match {
		return "", errors.New("incorrect username or password")
	}

	ud, qerr := s.userRepo.GetEmployeeByUsername(ctx, users[0].Username)
	if qerr != nil {
		return "", qerr
	}
	tokenString, err := createToken(users[0], ud)

	users[0].LastLogin = time.Now()
	users[0].UpdatedBy = user.Username
	qerr = s.userRepo.UpdateLastLogin(ctx, users[0])
	if qerr != nil {
		return "", qerr
	}

	return tokenString, err
}

func (s service) GetUser(ctx context.Context, username string) (res *payload.User, err error) {
	users, qerr := s.userRepo.GetUserByUsername(ctx, username)
	if len(users) == 0 || qerr != nil {
		return nil, errors.New("user not found")
	}

	em, qerr := s.userRepo.GetEmployeeByUsername(ctx, username)
	if qerr != nil {
		return nil, qerr
	}

	resUser := &payload.User{
		User:     users[0],
		Employee: em,
	}

	return resUser, err
}
