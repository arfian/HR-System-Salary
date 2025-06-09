package handler

import (
	"hr-system-salary/internal/app/attendance/port"

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
	// var (
	// 	dataUser payload.User
	// )
	// if err := c.ShouldBind(&dataUser); err != nil {
	// 	helper.ResponseError(c, err)
	// 	return
	// }

	// validate := validator.New()
	// err := validate.Struct(dataUser)
	// if err != nil {
	// 	helper.ResponseError(c, err)
	// 	return
	// }

	// res, err := h.userService.Register(c.Request.Context(), dataUser.User, dataUser.Employee)
	// if err != nil {
	// 	helper.ResponseError(c, err)
	// 	return
	// }

	// helper.ResponseData(c, &helper.Response{
	// 	Message: "register successfully",
	// 	Data:    res,
	// })
}
