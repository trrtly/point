package db

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
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

	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetConnMaxLifetime(30 * time.Second)
	db.DB().SetMaxOpenConns(100)

	return &DB{db}, nil
}

func setupDatabase(cfg *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Name)
	return gorm.Open(cfg.Type, dsn)
}

// updateTimeStampForCreateCallback will set `CreatedAt`, `UpdatedAt` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := uint(time.Now().Unix())
		if createdAtField, ok := scope.FieldByName("CreatedAt"); ok {
			if createdAtField.IsBlank {
				createdAtField.Set(nowTime)
			}
		}

		if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
			if updatedAtField.IsBlank {
				updatedAtField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `UpdatedAt` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdatedAt", uint(time.Now().Unix()))
	}
}
