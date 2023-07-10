package databases

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type ConfigMysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func NewMysqlDB(dbname string, cm ConfigMysql) (*sql.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cm.User, cm.Password, cm.Host, cm.Port, dbname)

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		return &sql.DB{}, err
	}

	err = db.Ping()

	if err != nil {
		fmt.Printf("Error Mysql Ping: %s", err)
	}

	return db, nil

}
