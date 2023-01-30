package connection

import (
	"fmt"
	"indonesian-heroes/hero"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection(env map[string]string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env["username"],
		env["password"],
		env["host"],
		env["port"],
		env["database"],
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// migrate schema
	if err := db.AutoMigrate(&hero.Hero{}); err != nil {
		return db, err
	}

	return db, nil
}
