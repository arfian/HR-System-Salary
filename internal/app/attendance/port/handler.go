package port

import (
	"github.com/gin-gonic/gin"
)

type IAttendanceHandler interface {

	// (POST /attendance/employee)
	AddAttendanceEmployee(ctx *gin.Context)

	// (POST /attendance/admin)
	// AddAttendanceAdmin(ctx *gin.Context)
}
