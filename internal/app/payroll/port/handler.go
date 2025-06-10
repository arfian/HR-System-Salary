package port

import (
	"github.com/gin-gonic/gin"
)

type IPayrollHandler interface {

	// (POST /generate-payroll)
	GeneratePayroll(ctx *gin.Context)
}
