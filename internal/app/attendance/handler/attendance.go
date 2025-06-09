package handler

import (
	"errors"
	"fmt"
	"hr-system-salary/internal/app/attendance/payload"
	"hr-system-salary/internal/app/attendance/port"
	"hr-system-salary/pkg/helper"
	"hr-system-salary/pkg/validations"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handler struct {
	attendanceService port.IAttendanceService
}

func New(attendanceService port.IAttendanceService) port.IAttendanceHandler {
	return &handler{
		attendanceService: attendanceService,
	}
}

func (h *handler) AddAttendanceEmployee(c *gin.Context) {
	username := c.GetString("username")
	res, err := h.attendanceService.AddAttendanceEmployee(c.Request.Context(), username)
	if err != nil {
		helper.ResponseError(c, err)
		return
	}
	status := ""
	if res != nil {
		status = res.Status
	}
	helper.ResponseData(c, &helper.Response{
		Message: fmt.Sprintf("insert %s successfully", status),
		Data:    res,
	})
}

func (h *handler) AddAttendanceAdmin(c *gin.Context) {
	username := c.GetString("username")
	rolename := c.GetString("rolename")

	var (
		paramAttendance payload.ParamBulkAttendance
	)

	if rolename != "admin" {
		helper.ResponseError(c, errors.New("only admins with access"))
		return
	}

	if err := c.ShouldBind(&paramAttendance); err != nil {
		helper.ResponseError(c, err)
		return
	}

	validate := validator.New()
	err := validate.Struct(paramAttendance)
	if err != nil {
		helper.ResponseError(c, err)
		return
	}

	startDate, err := time.Parse("2006-01-02", paramAttendance.StartDate)
	endDate, err := time.Parse("2006-01-02", paramAttendance.EndDate)
	if !validations.IsSameDateMonth(startDate, endDate) {
		helper.ResponseError(c, errors.New("date period must be the same month"))
		return
	}

	err = h.attendanceService.BulkInserAttendance(c.Request.Context(), paramAttendance, username)
	if err != nil {
		helper.ResponseError(c, err)
		return
	}

	helper.ResponseData(c, &helper.Response{
		Message: "insert successfully",
	})
}
