package config

import (
	"database/sql"
	"fmt"
	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var FccReadDB *gorm.DB
var FccWriteDB *gorm.DB

func openConn(dbConfig *DBConfig, maxOpenConns int, maxIdleConns int) *sql.DB {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?multiStatements=true",
			dbConfig.Username,
			dbConfig.Password,
			dbConfig.Host,
			dbConfig.Port,
			dbConfig.Database,
		))

	if err != nil {
		panic(err)
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	return db
}


func openDB(conn *sql.DB) *gorm.DB {
	db, err := gorm.Open(gorm_mysql.New(gorm_mysql.Config{Conn: conn}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func InitDb() {
	FccReadDB = openDB(openConn(&Global.DB.FccRead, 100, 100))
	FccWriteDB = openDB(openConn(&Global.DB.FccWrite, 10, 10))
}
