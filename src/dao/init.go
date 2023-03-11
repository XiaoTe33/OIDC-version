package dao

import (
	"database/sql"
	"fmt"
	"winter-examination/src/conf"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDb() {
	fmt.Println()
	var err error
	Db, err = sql.Open(conf.Database, conf.MysqlDNS)
	if err != nil {
		fmt.Println("mysql init failed......")
		return
	}
	err = Db.Ping()
	if err != nil {
		fmt.Println("mysql init failed......")
		return
	}
	fmt.Println("mysql init success")
}
