package model

import (
	"time"

	"github.com/go-openapi/strfmt"
)

type PayrollModel struct {
	ID                    strfmt.UUID4 `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Employee              string       `json:"employee" validate:"required"`
	PayrollDate           time.Time    `json"payroll_date" validate:"required"`
	CountAbsence          int          `json:"count_absence"`
	TotalAttendence       int          `json:"total_attendence"`
	BasicSalary           float32      `json:"basic_salary"`
	CountOvertime         int          `json:"count_overtime"`
	OvertimeRateHours     float32      `json:"overtime_rate_hours"`
	TotalOvertime         float32      `json:"total_overtime"`
	TotalDeductionAbsence float32      `json:"total_deduction_absence"`
	TotalReimbursement    float32      `json:"total_reimbursement"`
	TotalTakeHomePay      float32      `json:"total_take_home_pay"`
	Status                string       `json:"status"`
	CreatedBy             string       `json:"created_by"`
	UpdatedBy             string       `json:"updated_by" gorm:"default:null"`
	CreatedAt             time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt             time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt             *time.Time   `json:"deleted_at" gorm:"default:null"`
}

func (u PayrollModel) TableName() string {
	return "payroll"
}

type SettingPayrollModel struct {
	ID                strfmt.UUID4 `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	EndCutOff         int          `json:"end_cutoff"`
	OvertimeRateHours float32      `json:"overtime_rate_hours"`
	CreatedBy         string       `json:"created_by"`
	UpdatedBy         string       `json:"updated_by" gorm:"default:null"`
	CreatedAt         time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt         time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt         *time.Time   `json:"deleted_at" gorm:"default:null"`
}

func (u SettingPayrollModel) TableName() string {
	return "setting_payroll"
}
