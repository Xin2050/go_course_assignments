package mysql_db

import (
	"database/sql"
	"fmt"
	"github.com/Xin2050/go_course_assignments/s1/logger"
	_ "github.com/go-sql-driver/mysql"

	"time"
)

const (
	mysql_username = "go"
	mysql_password = "go"
	mysql_host     = "localhost:3306"
	mysql_db       = "go_bookstore"
)

var (
	Client *sql.DB
)

func init() {

	// create connection to server
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		mysql_username, mysql_password, mysql_host, mysql_db)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		logger.Error("can't connect to mysql server", err)
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		logger.Error("can't connect to mysql server", err)
		panic(err)
	}
	Client.SetConnMaxLifetime(time.Minute * 3)
	Client.SetMaxOpenConns(10)
	Client.SetMaxIdleConns(10)

	logger.Info("database successfully configured")

}
