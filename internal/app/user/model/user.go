package model

import (
	"time"

	"github.com/go-openapi/strfmt"
)

type UserModel struct {
	ID        strfmt.UUID4 `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Username  string       `json:"username" validate:"required"`
	Password  string       `json:"password" validate:"required"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt time.Time    `json:"deleted_at"`
}

func (u UserModel) TableName() string {
	return "users"
}

type UserDetailModel struct {
	ID          strfmt.UUID4 `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	UserId      string       `json:"user_id"`
	FirstName   string       `json:"firstname" gorm:"column:firstname"`
	LastName    string       `json:"lastname" gorm:"column:lastname"`
	Gender      string       `json:"gender" validate:"required,oneof=MALE FEMALE"`
	City        string       `json:"city"`
	Description string       `json:"description" gorm:"default:null"`
	CreatedAt   time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt   time.Time    `json:"deleted_at"`
}

func (u UserDetailModel) TableName() string {
	return "user_detail"
}

type UserPreferenceModel struct {
	ID              strfmt.UUID4 `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	UserId          string       `json:"user_id"`
	PreferredGender string       `json:"preferred_gender" validate:"required,oneof=MALE FEMALE"`
	PreferredCity   string       `json:"preferred_city"`
	CreatedAt       time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt       time.Time    `json:"deleted_at"`
}

func (u UserPreferenceModel) TableName() string {
	return "user_preference"
}

type UserStatisticModel struct {
	UserId string `json:"user_id"`
	Like   int    `json:"like"`
	Pass   string `json:"pass"`
}
