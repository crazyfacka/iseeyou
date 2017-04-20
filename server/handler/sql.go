package handler

import (
	"database/sql"
	"strconv"

	"github.com/crazyfacka/iseeyou/server/commons"
	_ "github.com/go-sql-driver/mysql" // SQL Driver
)

// SQL core struct
type SQL struct {
	db *sql.DB
}

var sqlh *SQL

// StorePing stores a ping message
func (s *SQL) StorePing() {

}

// GetSQLHandler intatiates this handler
func GetSQLHandler(cfg commons.Config) *SQL {
	var err error

	sqlh := &SQL{}

	sqlConf := cfg.SQL
	dsn := sqlConf.User + ":" + sqlConf.Password + "@" + sqlConf.Host + ":" + strconv.FormatInt(sqlConf.Port, 10) + "/" + sqlConf.DBName
	sqlh.db, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	commons.Debug("[SQL] Connected to %s @ %s:%d", sqlConf.DBName, sqlConf.Host, sqlConf.Port)

	return sqlh
}
