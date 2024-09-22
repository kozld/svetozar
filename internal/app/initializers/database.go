package initializers

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/kozld/svetozar/config"
)

func InitializeDatabase(conf *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(getConnString(conf)))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getConnString(conf *config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		conf.DBUser,
		conf.DBPwd,
		conf.DBHost,
		conf.DBPort,
		conf.DBName,
	)
}
