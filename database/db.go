package database

import (
	"fmt"
	"log"

	"github.com/vctrthe/api-go/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
	dialect := mysql.Open(getDSN())

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	db, err := gorm.Open(dialect, gormConfig)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

func getDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		config.C.Database.Username,
		config.C.Database.Password,
		config.C.Database.Host,
		config.C.Database.Port,
		config.C.Database.DBName,
		config.C.Database.Charset,
		config.C.Database.ParseTime,
		config.C.Database.Loc,
	)
}
