package db

import (
	"database/sql"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"hr-system-salary/config"
)

type GormDB struct {
	*gorm.DB
}

type dbConfig struct {
	GormDB       *GormDB
	ConnectionDB *sql.DB
}

func (db dbConfig) CloseConnection() error {
	return db.ConnectionDB.Close()
}

func Init(dsn string) (dbConfig, error) {
	var (
		dbConfigVar dbConfig
		loggerGorm  logger.Interface
	)
	configData := config.GetConfig()

	loggerGorm = logger.Default.LogMode(logger.Silent)
	if configData.App.Env == "local" {
		loggerGorm = logger.Default.LogMode(logger.Info)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger:                 loggerGorm,
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return dbConfigVar, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return dbConfigVar, err
	}

	sqlDB.SetConnMaxIdleTime(time.Second * time.Duration(configData.DB.MaxIdletimeConn))
	sqlDB.SetMaxIdleConns(configData.DB.MaxIdleConn)
	sqlDB.SetMaxOpenConns(configData.DB.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(configData.DB.MaxLifetimeConn))
	dbConfigVar.ConnectionDB = sqlDB

	dbConfigVar.GormDB = &GormDB{gormDB}
	log.Println("database connected")

	return dbConfigVar, nil
}
