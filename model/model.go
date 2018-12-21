package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type Employee struct {
	gorm.Model
	Name string `gorm:"unique" json:"name"`
	City string `json:"city"`
	Age string `json:"age"`
	Status bool `json:"status"`
}

func (e *Employee) Disable() {
	e.Status = false
}

func (e *Employee) Enable() {
	e.Status = true
}

// DBMigrate将会创建表，若有必要，还会创建某些关系。
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Employee{})
	return db
}
