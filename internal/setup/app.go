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
)

type InternalAppStruct struct {
	Repositories initRepositoriesApp
	Services     initServicesApp
	Handler      InitHandlerApp
}

type initRepositoriesApp struct {
	userRepo   userPorts.IUserRepository
	TrxHandler transaction.ISqlTransaction
	// HealthCheckRepo healthCheckPorts.IHealthCheckRepository
	dbInstance *gorm.DB
	// cache      cache.ICache
}

func initAppRepo(gormDB *db.GormDB, initializeApp *InternalAppStruct) {
	initializeApp.Repositories.userRepo = userRepo.NewRepository(gormDB)
	// initializeApp.Repositories.HealthCheckRepo = healthCheckRepo.NewHealthCheckRepository(gormDB.DB, rc)

	// Initiate trxRepo handler
	initializeApp.Repositories.TrxHandler = transaction.NewSqlTransaction(gormDB)

	// Get Gorm instance
	initializeApp.Repositories.dbInstance = gormDB.DB
}

type initServicesApp struct {
	UserService userPorts.IUserService
	// HealthCheckService healthCheckPorts.IHealthCheckService
}

func initAppService(initializeApp *InternalAppStruct) {
	// initializeApp.Services.HealthCheckService = healthCheckService.NewService(initializeApp.Repositories.HealthCheckRepo)
	initializeApp.Services.UserService = userService.New(initializeApp.Repositories.userRepo)
}

// HANDLER INIT
type InitHandlerApp struct {
	UserHandler userPorts.IUserHandler
	// HealthCheckHandler healthCheckPorts.IHealthCheckHandler
}

func initAppHandler(initializeApp *InternalAppStruct) {
	// initializeApp.Handler.HealthCheckHandler = healthCheckHandler.NewHealthCheckHandler(initializeApp.Services.HealthCheckService)
	initializeApp.Handler.UserHandler = userHandler.New(initializeApp.Services.UserService)
}
