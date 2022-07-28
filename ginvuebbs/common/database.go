package common

import (
	"database/sql"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

var (
	DB  *sql.DB
	err error
)

func InitDB() {
	var (
		driverName = viper.GetString("dataSourse.driverName")
		host       = viper.GetString("dataSourse.host")
		port       = viper.GetString("dataSourse.port")
		database   = viper.GetString("dataSourse.database")
		username   = viper.GetString("dataSourse.username")
		password   = viper.GetString("dataSourse.password")
		charset    = viper.GetString("dataSourse.charset")
	)
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	DB, err = sql.Open(driverName, args)
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	log.Println("连接成功")
}
