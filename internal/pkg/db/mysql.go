package db

import (
	"crud-go-server/internal/pkg/repo"
	"crud-go-server/internal/setting"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySqlSession(config setting.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("root:%s@tcp(%s:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		config.MYSQL_PASS,
		config.MYSQL_HOST,
		config.MYSQL_PORT,
		config.MYSQL_DB)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("init mysql session failed: %v", err.Error())
		return nil, err
	}

	db.DisableForeignKeyConstraintWhenMigrating = true

	err = db.AutoMigrate(
		&repo.StudentDao{},
		&repo.MarkDao{},
		&repo.ModuleDao{},
	)

	if err != nil {
		log.Fatalf(err.Error())
	}
	return db, nil
}
