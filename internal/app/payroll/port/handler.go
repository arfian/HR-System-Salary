package port

import (
	"github.com/gin-gonic/gin"
)

type IPayrollHandler interface {

	// (POST / )
	GeneratePayroll(ctx *gin.Context)

	// (POST /get-payslip)
	GetPayslip(ctx *gin.Context)

	// (POST /get-all-payslip)
	GetAllPayslip(ctx *gin.Context)
}
