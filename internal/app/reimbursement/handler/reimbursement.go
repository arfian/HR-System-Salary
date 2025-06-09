package handler

import (
	"hr-system-salary/internal/app/reimbursement/payload"
	"hr-system-salary/internal/app/reimbursement/port"
	"hr-system-salary/pkg/helper"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handler struct {
	reimbursementService port.IReimbursementService
}

func New(reimbursementService port.IReimbursementService) port.IReimbursementHandler {
	return &handler{
		reimbursementService: reimbursementService,
	}
}

func (h *handler) AddReimbursement(c *gin.Context) {
	username := c.GetString("username")
	var (
		paramReimbursement payload.ParamReimbursement
	)

	if err := c.ShouldBind(&paramReimbursement); err != nil {
		helper.ResponseError(c, err)
		return
	}

	validate := validator.New()
	err := validate.Struct(paramReimbursement)
	if err != nil {
		helper.ResponseError(c, err)
		return
	}

	res, err := h.reimbursementService.AddReimbursement(c.Request.Context(), username, paramReimbursement)
	if err != nil {
		helper.ResponseError(c, err)
		return
	}

	helper.ResponseData(c, &helper.Response{
		Message: "insert successfully",
		Data:    res,
	})
}
