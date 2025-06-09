package model

import (
	"time"

	"github.com/go-openapi/strfmt"
)

type ReimbursementModel struct {
	ID                  strfmt.UUID4 `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Employee            string       `json:"employee" validate:"required"`
	ReimbursementDate   time.Time    `json:"reimbursement_date" validate:"required"`
	ReimbursementType   string       `json:"reimbursement_type"`
	ReimbursementAmount float32      `json:"reimbursement_amount" validate:"required"`
	Description         string       `json:"description"`
	Status              string       `json:"status"`
	Payroll             string       `json:"payroll" gorm:"default:null"`
	CreatedBy           string       `json:"created_by"`
	UpdatedBy           string       `json:"updated_by" gorm:"default:null"`
	CreatedAt           time.Time    `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt           time.Time    `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt           *time.Time   `json:"deleted_at" gorm:"default:null"`
}

func (u ReimbursementModel) TableName() string {
	return "reimbursement"
}
