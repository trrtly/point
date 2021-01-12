package db

import (
	"fmt"
	"time"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

type (
	Config struct {
		Type     string `envconfig:"POINT_DATABASE_TYPE" default:"mysql"`
		User     string `envconfig:"POINT_DATABASE_USER"`
		Password string `envconfig:"POINT_DATABASE_PASSWORD"`
		Host     string `envconfig:"POINT_DATABASE_HOST"`
		Name     string `envconfig:"POINT_DATABASE_NAME"`
	}

	DB struct {
		*gorm.DB
	}
)

// Connect to a database.
func Connect(cfg *Config) (*DB, error) {
	db, err := setupDatabase(cfg)
	if err != nil {
		return &DB{db}, err
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(30 * time.Second)
	sqlDB.SetMaxOpenConns(100)

	return &DB{db}, nil
}

func setupDatabase(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Name)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
