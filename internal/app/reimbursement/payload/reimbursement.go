package payload

type ParamReimbursement struct {
	ReimbursementDate   string  `json:"reimbursement_date" validate:"required,datetime=2006-01-02"`
	Description         string  `json:"description"`
	ReimbursementAmount float32 `json:"reimbursement_amount" validate:"required"`
	ReimbursementType   string  `json:"reimbursement_type" validate:"required,gte=0"`
}
