package payload

import "time"

type ParamBulkAttendance struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
