package handler

import (
	"errors"
	"hr-system-salary/internal/app/payroll/payload"
	"hr-system-salary/internal/app/payroll/port"
	"hr-system-salary/pkg/helper"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handler struct {
	payrollService port.IPayrollService
}

func New(payrollService port.IPayrollService) port.IPayrollHandler {
	return &handler{
		payrollService: payrollService,
	}
}

func (h *handler) GeneratePayroll(c *gin.Context) {
	username := c.GetString("username")
	rolename := c.GetString("rolename")

	if rolename != "admin" {
		helper.ResponseError(c, errors.New("only admins with access"))
		return
	}

	var (
		paramGeneratePayroll payload.ParamGeneratePayroll
	)

	if err := c.ShouldBind(&paramGeneratePayroll); err != nil {
		helper.ResponseError(c, err)
		return
	}

	validate := validator.New()
	err := validate.Struct(paramGeneratePayroll)
	if err != nil {
		helper.ResponseError(c, err)
		return
	}

	err = h.payrollService.InsertPayroll(c.Request.Context(), paramGeneratePayroll, username)
	if err != nil {
		helper.ResponseError(c, err)
		return
	}

	helper.ResponseData(c, &helper.Response{
		Message: "insert successfully",
	})
}

func (h *handler) GetPayslip(c *gin.Context) {
	username := c.GetString("username")

	var (
		paramGeneratePayroll payload.ParamGeneratePayroll
	)

	if err := c.ShouldBind(&paramGeneratePayroll); err != nil {
		helper.ResponseError(c, err)
		return
	}

	validate := validator.New()
	err := validate.Struct(paramGeneratePayroll)
	if err != nil {
		helper.ResponseError(c, err)
		return
	}

	payrollData, err := h.payrollService.GetPayrollByMonth(c.Request.Context(), paramGeneratePayroll, username)
	helper.ResponseData(c, &helper.Response{
		Message: "get data successfully",
		Data:    payrollData,
	})
}

func (h *handler) GetAllPayslip(c *gin.Context) {
	username := c.GetString("username")
	rolename := c.GetString("rolename")

	if rolename != "admin" {
		helper.ResponseError(c, errors.New("only admins with access"))
		return
	}

	var (
		paramGeneratePayroll payload.ParamGeneratePayroll
	)

	if err := c.ShouldBind(&paramGeneratePayroll); err != nil {
		helper.ResponseError(c, err)
		return
	}

	validate := validator.New()
	err := validate.Struct(paramGeneratePayroll)
	if err != nil {
		helper.ResponseError(c, err)
		return
	}

	payrollData, err := h.payrollService.GetAllPayrollByMonth(c.Request.Context(), paramGeneratePayroll, username)
	helper.ResponseData(c, &helper.Response{
		Message: "get data successfully",
		Data:    payrollData,
	})
}
