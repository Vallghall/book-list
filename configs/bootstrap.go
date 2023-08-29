package configs

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/Vallghall/book-list/pkg/models"
	"github.com/Vallghall/book-list/pkg/store"
	"github.com/Vallghall/book-list/pkg/store/postgres"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	configPath = "configs/config.yaml"
)

var (
	// disableMigration - flag that disables gorm auto-migration
	// for bootstrap performance boost
	disableMigration bool
)

func init() {
	flag.BoolVar(&disableMigration, "no-migration", false, "Disables gorm auto-migration for bootstrap performance boost")
	flag.Parse()
}

// DBConf - wrapper for database configuration
type DBConf struct {
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SSLOpt   string `yaml:"ssl"`
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
	Port              string `yaml:"port"`
	DBLogLevel        int    `yaml:"db-log-level"`
	HLogLevel         string `yaml:"h-log-level"`
	SigningKey        string `yaml:"signing-key"`
	TemplatePath      string `yaml:"template-path"`
	TemplateExtension string `yaml:"template-extension"`
}

// Conf - represents whole project's configs
type Conf struct {
	*DBConf  `yaml:"db"`
	*AppConf `yaml:"app"`
	db       *gorm.DB
	logger   *zap.Logger
}

// Store - Repository instance getter
func (c *Conf) Store() store.Store {
	return postgres.New(c.db)
}

// LogLevel - validates log level parsed from congigs
// and returns it if it is within [1;4] range
func (c *Conf) LogLevel() (logger.LogLevel, error) {
	if c.DBLogLevel < 0 || c.DBLogLevel > 4 {
		return 0, fmt.Errorf("invalid log level: %d", c.DBLogLevel)
	}

	return logger.LogLevel(c.DBLogLevel), nil
}

// HandlerLogLevel - logger getter
func (c *Conf) HandlerLogLevel() *zap.Logger {
	return c.logger
}

// Bootstrap - bootstraps the application loading migrations etc
func Bootstrap() (*Conf, error) {
	var err error
	time.Local, err = time.LoadLocation("Europe/Moscow")
	if err != nil {
		return nil, err
	}

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

	lvl, err := c.LogLevel()
	if err != nil {
		return nil, err
	}
	c.db, err = gorm.Open(
		pg.Open(c.ConnString()),
		&gorm.Config{
			Logger: logger.Default.LogMode(lvl),
		})
	if err != nil {
		return nil, err
	}

	if !disableMigration {
		err = c.db.AutoMigrate(&models.User{}, &models.Author{}, &models.Book{})
		if err != nil {
			return nil, err
		}
	}

	hLevel, err := zapcore.ParseLevel(c.HLogLevel)
	if err != nil {
		return nil, err
	}

	lc := zap.NewProductionConfig()
	lc.Level.SetLevel(hLevel)
	c.logger, err = lc.Build()
	if err != nil {
		return nil, err
	}

	return c, nil
}
