package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	mysql_username = os.Getenv("DB_USERNAME")
	mysql_password = os.Getenv("DB_PASSWORD")
	mysql_host     = os.Getenv("DB_HOST")
	mysql_port     = os.Getenv("DB_PORT")
	mysql_schema   = os.Getenv("DB_NAME")
)

var (
	Client *sql.DB
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		mysql_username, mysql_password, mysql_host, mysql_port, mysql_schema,
	)

	log.Println(fmt.Sprintf("about to connect to %s", dataSourceName))

	var err error
	Client, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")

}
