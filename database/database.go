package database

import (
	"database/sql"

	"github.com/yasrilimam98/grb-restapi/config"
)

var db *sql.DB
var err error

func Init() {
	conf := config.GetConfig()
	connetionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	db, err = sql.Open("mysql", connetionString)

	if err != nil {
		panic("Connection String error!")
	}

	err = db.Ping()
	if err != nil {
		panic("DSN Invalid")
	}
}

func CreateCon() *sql.DB {
	return db
}
