package rest

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"hr-system-salary/pkg/constants"
	"hr-system-salary/pkg/validations"

	"hr-system-salary/cmd/rest/middleware"
	"hr-system-salary/config"

	"hr-system-salary/internal/setup"

	attendanceServer "hr-system-salary/internal/app/attendance/server"
	reimbursementServer "hr-system-salary/internal/app/reimbursement/server"
	userServer "hr-system-salary/internal/app/user/server"

	_ "hr-system-salary/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func StartServer(setupData setup.SetupData) {
	conf := config.GetConfig()
	if conf.App.Env == constants.PRODUCTION {
		gin.SetMode(gin.ReleaseMode)
	}

	// GIN Init
	router := gin.Default()
	router.UseRawPath = true
	validations.InitStructValidation()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// router.GET("/health", setupData.InternalApp.Handler.HealthCheckHandler.HealthCheck)

	router.Use(middleware.CORSMiddleware())

	initPublicRoute(router, setupData.InternalApp)

	router.Use(middleware.JWTAuthMiddleware())

	initRoute(router, setupData.InternalApp)

	port := config.GetConfig().Http.Port
	httpServer := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: router,
	}

	go func() {
		// service connections
		if err := httpServer.ListenAndServe(); err != nil {
			log.Println("listen:", err)
		}
	}()
	log.Println("webserver started")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}

	_ = setup.CloseDB()

	log.Println("Server exiting")
}

func initRoute(router *gin.Engine, internalAppStruct setup.InternalAppStruct) {
	apiRouter := router.Group("/v1/api")
	userServer.Routes.NewProfile(apiRouter.Group("/profile"), internalAppStruct.Handler.UserHandler)
	attendanceServer.Routes.New(apiRouter.Group("/attendance"), internalAppStruct.Handler.AttendanceHandler)
	reimbursementServer.Routes.New(apiRouter.Group("/reimbursement"), internalAppStruct.Handler.ReimbursementHandler)
}

func initPublicRoute(router *gin.Engine, internalAppStruct setup.InternalAppStruct) {
	apiRouter := router.Group("/v1/public-api")

	userServer.Routes.New(apiRouter.Group("/user"), internalAppStruct.Handler.UserHandler)
}
