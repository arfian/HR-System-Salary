package server

import (
	"github.com/gin-gonic/gin"

	"hr-system-salary/internal/app/attendance/port"
)

type (
	routes struct{}
)

var (
	Routes routes
)

func (r routes) New(router *gin.RouterGroup, handler port.IAttendanceHandler) {
	router.POST("/employee", handler.AddAttendanceEmployee)
	router.POST("/admin", handler.AddAttendanceAdmin)
	router.POST("/overtime", handler.AddOvertime)
}
