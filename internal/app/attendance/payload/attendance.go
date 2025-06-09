package payload

type ParamBulkAttendance struct {
	StartDate  string `json:"start_date" validate:"required,datetime=2006-01-02"`
	EndDate    string `json:"end_date" validate:"required,datetime=2006-01-02"`
	EmployeeID string `json:"employee_id" validate:"required"`
}
