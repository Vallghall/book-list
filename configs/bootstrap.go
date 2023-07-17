package configs

import (
	"embed"
	"fmt"
	"os"

	"github.com/Vallghall/book-list/pkg/store"
	"github.com/Vallghall/book-list/pkg/store/postgres"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"gopkg.in/yaml.v3"
)

const (
	configPath = "configs/config.yaml"
	driver     = "postgres"
)

// DBConf - wrapper for database configuration
type DBConf struct {
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SSLOpt   string `yaml:"ssl"`
	//SSlCert  string `yaml:"sslrootsert"`
}

// ConnString returns a connection string in a required format
func (c *DBConf) ConnString() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
		c.SSLOpt,
	)
}

// AppConf - wrapper for app configuration
type AppConf struct {
	Port string `yaml:"port"`
}

// Conf - represents whole project's configs
type Conf struct {
	*DBConf  `yaml:"db"`
	*AppConf `yaml:"app"`
	dbHandle *sqlx.DB `yaml:"-"`
}

func (c *Conf) Store() store.Store {
	return postgres.New(c.dbHandle)
}

// Bootstrap - bootstraps thr application loading migrations etc
func Bootstrap(migrations embed.FS) (*Conf, error) {
	f, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	c := new(Conf)

	err = yaml.NewDecoder(f).Decode(c)
	if err != nil {
		return nil, err
	}

	c.dbHandle, err = getConnection(c.DBConf)
	if err != nil {
		return nil, err
	}

	goose.SetBaseFS(migrations)
	err = goose.SetDialect(driver)
	if err != nil {
		return nil, err
	}

	err = goose.Up(c.dbHandle.DB, "script/migrations")
	if err != nil {
		return nil, err
	}

	return c, nil
}

// getConnection - returns a handle to a database
func getConnection(conf *DBConf) (*sqlx.DB, error) {
	return sqlx.Connect(driver, conf.ConnString())
}
