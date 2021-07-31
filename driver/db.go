package driver

import (
	"cpbl-data-parser/utils"
	_ "cpbl-data-parser/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var mySqlDB *gorm.DB

func DatabaseClient() *gorm.DB {
	dbUser := utils.GetEnv("DATABASE_USERNAME")
	dbPassword := utils.GetEnv("DATABASE_PASSWORD")
	dbHost := utils.GetEnv("DATABASE_HOST")
	dbSchema := utils.GetEnv("DATABASE_SCHEMA")
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True", dbUser, dbPassword, dbHost, dbSchema)

	var err error
	logger.Debugln("Database Info:", dsn)
	var logLevel gLogger.LogLevel
	env := utils.GetEnv("LOG_LEVEL")
	if env == "debug" {
		logLevel = gLogger.Info
	} else {
		logLevel = gLogger.Silent
	}

	newLogger := gLogger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), gLogger.Config{
		LogLevel:                  logLevel,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	})

	mySqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		logger.Panicln("Database connection error")
		return nil
	}
	sqlDB, err := mySqlDB.DB()

	if err != nil {
		logger.Panicln("Invalid database instance")
		return nil
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	return mySqlDB
}
