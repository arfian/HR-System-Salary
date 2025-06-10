package server

import (
	"github.com/gin-gonic/gin"

	"hr-system-salary/internal/app/payroll/port"
)

type (
	routes struct{}
)

var (
	Routes routes
)

func (r routes) New(router *gin.RouterGroup, handler port.IPayrollHandler) {
	router.POST("/", handler.GeneratePayroll)
}
