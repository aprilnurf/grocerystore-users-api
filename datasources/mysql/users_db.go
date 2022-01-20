package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host = "mysql_users_host"
	mysql_users_schema = "mysql_users_schema"
)

var (
	Client *sql.DB
	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host = os.Getenv(mysql_users_host)
	schema = os.Getenv(mysql_users_schema)
)

func init() {
	//rout@tcp(localhost:3606)/belajar
	datasourceName := fmt.Sprintf("%s@tcp(%s)/%s?charset=utf8",
		username, host, schema)
	var err error
	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	//Client.SetConnMaxIdleTime(5)
	//Client.SetMaxOpenConns(20)
	//Client.SetConnMaxLifetime(60 * time.Minute)
	//Client.SetConnMaxIdleTime(10 * time.Minute)

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	//mysql.SetLogger()
	log.Println("db successfully connected")
}
