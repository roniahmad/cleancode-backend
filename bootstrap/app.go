package bootstrap

import (
	"database/sql"

	"github.com/rs/zerolog"
	"wetees.com/domain"
)

type Application struct {
	Conf *domain.Config
	Db   *sql.DB
	Log  zerolog.Logger
}

func App() Application {
	app := &Application{}
	app.Conf = NewConf()
	app.Log = NewLog(app.Conf)
	app.Db = NewDB(app.Conf, app.Log)

	return *app
}

func (app *Application) CloseDBConnection() {
	CloseDBConn(app.Db, app.Log)
}
