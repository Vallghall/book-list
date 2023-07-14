package bootstrap

import (
	"github.com/jmoiron/sqlx"
)

const driver = "postgres"

func GetConnection(conf *Conf) (*sqlx.DB, error) {
	return sqlx.Connect(driver, conf.ConnString())
}
