package handler

import (
	"fmt"
	"hr-system-salary/internal/app/attendance/port"
	"hr-system-salary/pkg/helper"

	"github.com/gin-gonic/gin"
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
