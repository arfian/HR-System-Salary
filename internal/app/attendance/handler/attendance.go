package handler

import (
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
	err := h.attendanceService.AddAttendanceEmployee(c.Request.Context(), username)
	if err != nil {
		helper.ResponseError(c, err)
		return
	}

	helper.ResponseData(c, &helper.Response{
		Message: "insert attendance successfully",
	})
}
