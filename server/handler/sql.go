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
func (s *SQL) StorePing(timestamp float64) bool {
	stmt, err := s.db.Prepare("INSERT INTO alive(timestamp) VALUES(?)")
	if err != nil {
		commons.Debug("[SQL] Error preparing statement: %s", err.Error())
		return false
	}

	res, err := stmt.Exec(timestamp)
	if err != nil {
		commons.Debug("[SQL] Error executing statement: %s", err.Error())
		return false
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		commons.Debug("[SQL] Error getting last ID: %s", err.Error())
		return false
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		commons.Debug("[SQL] Error getting row count: %s", err.Error())
		return false
	}

	commons.Debug("[SQL] ID = %d, affected = %d\n", lastID, rowCnt)
	return true
}

// GetSQLHandler intatiates this handler
func GetSQLHandler(cfg commons.Config) *SQL {
	var err error

	sqlh := &SQL{}

	sqlConf := cfg.SQL
	dsn := sqlConf.User + ":" + sqlConf.Password + "@tcp(" + sqlConf.Host + ":" + strconv.FormatInt(sqlConf.Port, 10) + ")/" + sqlConf.DBName
	sqlh.db, err = sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	commons.Debug("[SQL] Connected to %s @ %s:%d", sqlConf.DBName, sqlConf.Host, sqlConf.Port)

	return sqlh
}
