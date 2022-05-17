package mysql

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Config defines the database configuration
type Config struct {
	Protocol        string        `yaml:"protocol" json:"protocol"`
	Host            string        `yaml:"host" json:"host"`
	User            string        `yaml:"user" json:"user"`
	Password        string        `yaml:"password" json:"password"`
	DB              string        `yaml:"db" json:"db"`
	ReadTimeout     time.Duration `yaml:"readTimeout" json:"readTimeout"`
	WriteTimeout    time.Duration `yaml:"writeTimeout" json:"writeTimeout"`
	ParseTime       bool          `yaml:"parseTime" json:"parseTime"`
	Loc             time.Location `yaml:"location" json:"loc"`
	MaxOpenConns    int           `yaml:"maxOpenConns" json:"maxOpenConns"`
	MaxIdleConns    int           `yaml:"maxIdleConns" json:"maxIdleConns"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime" json:"connMaxLifetime"`
	VerifyConnCheck bool          `yaml:"verifyConnCheck" json:"verifyConnCheck"`
}

// Connect connects to a mysql database and returns the connection
func Connect(cfg Config) (*sqlx.DB, error) {
	conf := mysql.NewConfig()
	conf.Net = cfg.Protocol
	conf.Addr = cfg.Host
	conf.User = cfg.User
	conf.Passwd = cfg.Password
	conf.DBName = cfg.DB
	conf.ParseTime = cfg.ParseTime
	conf.Loc = &cfg.Loc

	db, err := sqlx.Open("mysql", conf.FormatDSN())
	if err != nil {
		return db, errors.Wrap(err, "could not connect to database")
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	if err := db.Ping(); err != nil {
		return db, errors.Wrap(err, "could not ping database")
	}

	if cfg.VerifyConnCheck {
		if err := Check(db); err != nil {
			return db, errors.Wrap(err, "could not verify connectivity")
		}
	}

	return db, nil
}

// Check checks if the connection to the database is really established
// by querying the simplest query possible
func Check(db *sqlx.DB) error {
	const q = `SELECT true`
	var tmp bool
	return db.QueryRow(q).Scan(&tmp)
}
