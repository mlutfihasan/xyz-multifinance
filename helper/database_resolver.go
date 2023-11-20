package helper

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

func ConnectMysqlDatabaseResolver(configuration Configuration) *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	database, err := gorm.Open(
		mysql.Open(configuration.DbDsnSource1),
		&gorm.Config{Logger: newLogger},
	)
	if err != nil {
		panic("failed to connect database")
	}

	plugin := dbresolver.Register(dbresolver.Config{
		Replicas:          []gorm.Dialector{mysql.Open(configuration.DbDsnReplication1)},
		Policy:            dbresolver.RandomPolicy{},
		TraceResolverMode: true,
	})
	database.Use(plugin)

	return database
}
