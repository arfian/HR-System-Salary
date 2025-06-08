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

func (s *service) Register(ctx context.Context, user model.UserModel, ud model.UserDetailModel, up model.UserPreferenceModel) (token string, err error) {
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

	user.Password = hash
	user, qerr = s.userRepo.InsertUser(ctx, user)
	if qerr != nil {
		return "", qerr
	}

	ud.UserId = user.ID.String()
	qerr = s.userRepo.InsertUserDetail(ctx, ud)
	if qerr != nil {
		return "", qerr
	}

	up.UserId = user.ID.String()
	qerr = s.userRepo.InsertUserPreference(ctx, up)
	if qerr != nil {
		return "", qerr
	}
	tokenString, err := createToken(user, ud)

	return tokenString, err
}

func createToken(user model.UserModel, ud model.UserDetailModel) (string, error) {
	configData := config.GetConfig()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        user.ID,
		"username":  user.Username,
		"firstname": ud.FirstName,
		"lastname":  ud.LastName,
		"exp":       time.Now().Add(time.Hour).Unix(),
		"iat":       time.Now().Unix(),
	})
	tokenString, err := claims.SignedString([]byte(configData.JWT.SigningKey))

	return tokenString, err
}

func (s service) Login(ctx context.Context, user model.UserModel) (token string, err error) {
	users, qerr := s.userRepo.GetPasswordByUsername(ctx, user.Username)
	if len(users) == 0 || qerr != nil {
		return "", errors.New("incorrect username or password")
	}

	match := encrypt.CheckPasswordHash(user.Password, users[0].Password)
	if !match {
		return "", errors.New("incorrect username or password")
	}

	ud, qerr := s.userRepo.GetUserDetailById(ctx, users[0].ID.String())
	if qerr != nil {
		return "", qerr
	}
	tokenString, err := createToken(users[0], ud)

	return tokenString, err
}

func (s service) GetUser(ctx context.Context, username string) (res *payload.User, err error) {
	users, qerr := s.userRepo.GetUserByUsername(ctx, username)
	if len(users) == 0 || qerr != nil {
		return nil, errors.New("user not found")
	}

	ud, qerr := s.userRepo.GetUserDetailById(ctx, users[0].ID.String())
	if qerr != nil {
		return nil, qerr
	}

	up, qerr := s.userRepo.GetUserPreference(ctx, users[0].ID.String())
	if qerr != nil {
		return nil, qerr
	}
	resUser := &payload.User{
		User:           users[0],
		UserDetail:     ud,
		UserPreference: up,
	}

	return resUser, err
}
