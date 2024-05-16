package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	mysqlconnect "project-v/pkg/mysql"
)

func RunMysqlMigrate(cfg *mysqlconnect.Config) error {
	db := mysqlconnect.NewConnectMysql(cfg)
	driver, err := mysql.WithInstance(
		db, &mysql.Config{
			NoLock: true,
		},
	)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrate",
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
