package payload

import (
	model "hr-system-salary/internal/app/payroll/model"
	modelReimbursement "hr-system-salary/internal/app/reimbursement/model"
)

type ParamGeneratePayroll struct {
	PayrollDate string `json:"payroll_date" validate:"required,datetime=2006-01"`
}

type ResPayslip struct {
	Payslip        model.PayrollModel                      `json:"payslip"`
	Reimbursements []modelReimbursement.ReimbursementModel `json:"reimbursements"`
}

type ResAllPayslip struct {
	Payslip          []model.PayrollModel `json:"payslip"`
	TotalTakeHomePay float32              `json:"total_all_take_home_pay"`
}
