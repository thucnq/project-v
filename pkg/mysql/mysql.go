package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Host     string `json:"host"     mapstructure:"host"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Database string `json:"database" mapstructure:"database"`

	Port        int `json:"port"          mapstructure:"port"`
	MaxOpenCon  int `json:"max_open_con"  mapstructure:"max_open_con"`
	MaxIdleCon  int `json:"max_idle_con"  mapstructure:"max_idle_con"`
	MaxLifeTime int `json:"max_life_time" mapstructure:"max_life_time"`
}

func (cfg *Config) Prepare() {
	if cfg.Host == "" {
		panic("mysql need host")
	}

	if cfg.Username == "" {
		panic("mysql need username")
	}

	if cfg.Password == "" {
		panic("mysql need password")
	}

	if cfg.Port == 0 {
		cfg.Port = 3306
	}

	if cfg.Database == "" {
		panic("mysql need database")
	}

	if cfg.MaxOpenCon == 0 {
		cfg.MaxOpenCon = 5
	}

	if cfg.MaxIdleCon == 0 {
		cfg.MaxIdleCon = 5
	}

	if cfg.MaxLifeTime == 0 {
		cfg.MaxIdleCon = 5
	}

}

func (cfg Config) GetUri() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
}

func NewConnectMysql(cfg *Config) *sql.DB {
	var err error
	cfg.Prepare()
	uri := cfg.GetUri()
	db, err := sql.Open("mysql", uri)
	if err != nil {
		panic(err)
	}

	// https://www.alexedwards.net/blog/configuring-sqldb
	db.SetMaxOpenConns(cfg.MaxOpenCon)
	db.SetMaxIdleConns(cfg.MaxIdleCon)
	db.SetConnMaxLifetime(time.Duration(cfg.MaxLifeTime) * time.Minute)

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}
