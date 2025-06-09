package model

import (
	"time"

	"github.com/go-openapi/strfmt"
)

type RoleUserModel struct {
	ID        strfmt.UUID4 `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Rolename  string       `json:"rolename" validate:"required"`
	CreatedBy string       `json:"created_by" validate:"required"`
	UpdatedBy string       `json:"updated_by"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt time.Time    `json:"deleted_at" gorm:"default:null"`
}

func (u RoleUserModel) TableName() string {
	return "role_user"
}

type AuthUserModel struct {
	ID        strfmt.UUID4 `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Username  string       `json:"username" validate:"required"`
	Password  string       `json:"password" validate:"required"`
	IsActive  bool         `json:"is_active"`
	LastLogin time.Time    `json:"last_login"`
	CreatedBy string       `json:"created_by"`
	UpdatedBy string       `json:"updated_by"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt time.Time    `json:"deleted_at" gorm:"default:null"`
}

func (u AuthUserModel) TableName() string {
	return "auth_user"
}

type EmployeeModel struct {
	ID           strfmt.UUID4 `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Username     string       `json:"username"`
	Fullname     string       `json:"fullname"`
	Rolename     string       `json:"rolename"`
	Gender       string       `json:"gender" validate:"required,oneof=MALE FEMALE"`
	DateJoin     string       `json:"date_join"`
	SalaryAmount float32      `json:"salary_amount" gorm:"type:numeric(10,2);not null"`
	CreatedBy    string       `json:"created_by"`
	UpdatedBy    string       `json:"updated_by"`
	CreatedAt    time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt    time.Time    `json:"deleted_at" gorm:"default:null"`
}

func (u EmployeeModel) TableName() string {
	return "employee"
}
