package lib

import (
	"log"
	"os"
	"time"

	"github.com/wildanie12/region-api/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ConnectMySQL establish connection to main mysql database.
func (di *DatabaseInstance) ConnectMySQL() {
	var err error
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: config.Get().MySQL.ConnectionString,
	}), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             2 * time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Warn,
		}),
	})
	if err != nil {
		panic(err)
	}
	di.MySQL = db
}
