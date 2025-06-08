package payload

import (
	"hr-system-salary/internal/app/user/model"
)

type User struct {
	User           model.UserModel           `json:"user"`
	UserDetail     model.UserDetailModel     `json:"user_detail"`
	UserPreference model.UserPreferenceModel `json:"user_preference"`
}
