package bootstrap

import "fmt"

type DBConf struct {
	Host     string
	Name     string
	Port     string
	User     string
	Password string
}

func (c *DBConf) ConnString() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Name,
		"disabled",
	)
}

type Conf struct {
	DBConf
}
