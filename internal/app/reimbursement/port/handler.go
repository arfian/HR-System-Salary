package port

import (
	"github.com/gin-gonic/gin"
)

type IReimbursementHandler interface {

	// (POST /reimbursement)
	AddReimbursement(ctx *gin.Context)
}
