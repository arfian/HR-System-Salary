package model

import (
	"time"

	"github.com/go-openapi/strfmt"
)

type AttendanceModel struct {
	ID        strfmt.UUID4 `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Employee  string       `json:"employee" validate:"required"`
	CheckIn   time.Time    `json:"check_in"`
	CheckOut  time.Time    `json:"check_out" gorm:"default:null"`
	Status    string       `json:"status"`
	CreatedBy string       `json:"created_by"`
	UpdatedBy string       `json:"updated_by"`
	CreatedAt time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt *time.Time   `json:"deleted_at" gorm:"default:null"`
}

func (u AttendanceModel) TableName() string {
	return "attendance"
}

type OvertimeModel struct {
	ID            strfmt.UUID4 `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Employee      string       `json:"employee" validate:"required"`
	OvertimeHours int          `json:"overtime_hours" validate:"required"`
	OvertimeDate  time.Time    `json:"overtime_date" validate:"required"`
	Status        string       `json:"status"`
	Payroll       string       `json:"payroll" gorm:"default:null"`
	CreatedBy     string       `json:"created_by"`
	UpdatedBy     string       `json:"updated_by"`
	CreatedAt     time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt     time.Time    `json:"deleted_at" gorm:"default:null"`
}

func (u OvertimeModel) TableName() string {
	return "overtime"
}
