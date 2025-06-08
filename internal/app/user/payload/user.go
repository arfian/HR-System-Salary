package payload

import (
	"hr-system-salary/internal/app/user/model"
)

type User struct {
	User     model.AuthUserModel `json:"auth_user"`
	Employee model.EmployeeModel `json:"employee"`
}
