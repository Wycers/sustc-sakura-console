package service

import (
	"os"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/wycers/sustc-sakura-console/log"
	"github.com/wycers/sustc-sakura-console/util"
)

var logger = log.NewLogger(os.Stdout)

var db *gorm.DB

func ConnectDB() {
	var err error
	db, err = gorm.Open("sqlite3", util.Config.SQLite)
	if err != nil {
		logger.Fatalf("opens database failed: " + err.Error())
	}
	logger.Debug("used [SQLite] as underlying database")

	if err = db.AutoMigrate(util.Models...).Error; err != nil {
		logger.Fatal("auto migrate tables failed: " + err.Error())
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
}

func DisconnectDB() {
	if err := db.Close(); nil != err {
		logger.Errorf("Disconnect from database failed: " + err.Error())
	}
}

