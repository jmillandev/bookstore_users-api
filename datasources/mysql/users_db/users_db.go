package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmillandev/bookstore_utils-go/logger"
	"github.com/joho/godotenv"
)

var (
	mysql_schema, mysql_host, mysql_password, mysql_username, mysql_port string
)

var (
	Client *sql.DB
)

func set_mysql_config() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	mysql_username = os.Getenv("MYSQL_USER")
	mysql_password = os.Getenv("MYSQL_PASSWORD")
	mysql_host = os.Getenv("MWSQL_HOST")
	mysql_port = os.Getenv("MYSQL_PORT")
	mysql_schema = os.Getenv("MYSQL_DATABASE")
}

func init() {
	set_mysql_config()
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		mysql_username, mysql_password, mysql_host, mysql_port, mysql_schema,
	)

	log.Println(fmt.Sprintf("about to connect to %s", dataSourceName))

	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	mysql.SetLogger(logger.GetLogger())
	log.Println("database successfully configured")

}
