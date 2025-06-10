package payload

type ParamGeneratePayroll struct {
	PayrollDate string `json:"payroll_date" validate:"required,datetime=2006-01"`
}
