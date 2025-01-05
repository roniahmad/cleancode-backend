package bootstrap

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	"wetees.com/domain"
)

const DbDriver = "mysql"

func NewDB(conf *domain.Config, logger zerolog.Logger) *sql.DB {
	dsn := connString(conf)

	dbConn, err := sql.Open(DbDriver, dsn)
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}

	err = dbConn.Ping()
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}

	return dbConn
}

func connString(conf *domain.Config) string {
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		conf.Db.User, conf.Db.Password, conf.Db.Host, conf.Db.Port, conf.Db.DbName, val.Encode())

	return conn
}

func CloseDBConn(dbConn *sql.DB, logger zerolog.Logger) {
	if dbConn == nil {
		return
	}

	err := dbConn.Close()
	if err != nil {
		logger.Fatal().Msg(err.Error())
	}

	logger.Info().Msg("Connection to Database closed.")
}
