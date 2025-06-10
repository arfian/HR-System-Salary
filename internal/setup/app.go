package setup

import (
	"gorm.io/gorm"

	"hr-system-salary/pkg/transaction"

	"hr-system-salary/config/db"

	// healthCheckHandler "hr-system-salary/internal/app/healthcheck/handler"
	// healthCheckPorts "hr-system-salary/internal/app/healthcheck/port"
	// healthCheckRepo "hr-system-salary/internal/app/healthcheck/repository"
	// healthCheckService "hr-system-salary/internal/app/healthcheck/service"

	userHandler "hr-system-salary/internal/app/user/handler"
	userPorts "hr-system-salary/internal/app/user/port"
	userRepo "hr-system-salary/internal/app/user/repository"
	userService "hr-system-salary/internal/app/user/service"

	attendanceHandler "hr-system-salary/internal/app/attendance/handler"
	attendancePorts "hr-system-salary/internal/app/attendance/port"
	attendanceRepo "hr-system-salary/internal/app/attendance/repository"
	attendanceService "hr-system-salary/internal/app/attendance/service"

	reimbursementHandler "hr-system-salary/internal/app/reimbursement/handler"
	reimbursementPorts "hr-system-salary/internal/app/reimbursement/port"
	reimbursementRepo "hr-system-salary/internal/app/reimbursement/repository"
	reimbursementService "hr-system-salary/internal/app/reimbursement/service"

	payrollHandler "hr-system-salary/internal/app/payroll/handler"
	payrollPorts "hr-system-salary/internal/app/payroll/port"
	payrollRepo "hr-system-salary/internal/app/payroll/repository"
	payrollService "hr-system-salary/internal/app/payroll/service"
)

type InternalAppStruct struct {
	Repositories initRepositoriesApp
	Services     initServicesApp
	Handler      InitHandlerApp
}

type initRepositoriesApp struct {
	userRepo          userPorts.IUserRepository
	attendanceRepo    attendancePorts.IAttendanceRepository
	reimbursementRepo reimbursementPorts.IReimbursementRepository
	payrollRepo       payrollPorts.IPayrollRepository
	TrxHandler        transaction.ISqlTransaction
	// HealthCheckRepo healthCheckPorts.IHealthCheckRepository
	dbInstance *gorm.DB
	// cache      cache.ICache
}

func initAppRepo(gormDB *db.GormDB, initializeApp *InternalAppStruct) {
	initializeApp.Repositories.userRepo = userRepo.NewRepository(gormDB)
	initializeApp.Repositories.attendanceRepo = attendanceRepo.NewRepository(gormDB)
	initializeApp.Repositories.reimbursementRepo = reimbursementRepo.NewRepository(gormDB)
	initializeApp.Repositories.payrollRepo = payrollRepo.NewRepository(gormDB)
	// initializeApp.Repositories.HealthCheckRepo = healthCheckRepo.NewHealthCheckRepository(gormDB.DB, rc)

	// Initiate trxRepo handler
	initializeApp.Repositories.TrxHandler = transaction.NewSqlTransaction(gormDB)

	// Get Gorm instance
	initializeApp.Repositories.dbInstance = gormDB.DB
}

type initServicesApp struct {
	UserService          userPorts.IUserService
	AttendanceService    attendancePorts.IAttendanceService
	ReimbursementService reimbursementPorts.IReimbursementService
	PayrollService       payrollPorts.IPayrollService
	// HealthCheckService healthCheckPorts.IHealthCheckService
}

func initAppService(initializeApp *InternalAppStruct) {
	// initializeApp.Services.HealthCheckService = healthCheckService.NewService(initializeApp.Repositories.HealthCheckRepo)
	initializeApp.Services.UserService = userService.New(initializeApp.Repositories.userRepo)
	initializeApp.Services.AttendanceService = attendanceService.New(initializeApp.Repositories.attendanceRepo, initializeApp.Repositories.userRepo)
	initializeApp.Services.ReimbursementService = reimbursementService.New(initializeApp.Repositories.reimbursementRepo, initializeApp.Repositories.userRepo)
	initializeApp.Services.PayrollService = payrollService.New(initializeApp.Repositories.payrollRepo, initializeApp.Repositories.userRepo)
}

// HANDLER INIT
type InitHandlerApp struct {
	UserHandler          userPorts.IUserHandler
	AttendanceHandler    attendancePorts.IAttendanceHandler
	ReimbursementHandler reimbursementPorts.IReimbursementHandler
	PayrollHandler       payrollPorts.IPayrollHandler
	// HealthCheckHandler healthCheckPorts.IHealthCheckHandler
}

func initAppHandler(initializeApp *InternalAppStruct) {
	// initializeApp.Handler.HealthCheckHandler = healthCheckHandler.NewHealthCheckHandler(initializeApp.Services.HealthCheckService)
	initializeApp.Handler.UserHandler = userHandler.New(initializeApp.Services.UserService)
	initializeApp.Handler.AttendanceHandler = attendanceHandler.New(initializeApp.Services.AttendanceService)
	initializeApp.Handler.ReimbursementHandler = reimbursementHandler.New(initializeApp.Services.ReimbursementService)
	initializeApp.Handler.PayrollHandler = payrollHandler.New(initializeApp.Services.PayrollService)
}
