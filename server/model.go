package main

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 公司
type Company struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
}

// 车队
type Team struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Company     Company `gorm:"foreignKey:ID;onDelete:CASCADE"`
	ManagerName string
}

// 路线
type Route struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Team Team `gorm:"foreignKey:ID;onDelete:CASCADE"`
}

// 司机
type Driver struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Route Route `gorm:"foreignKey:ID;onDelete:CASCADE"`
}

// 队长
type RoadManager struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Route Route `gorm:"foreignKey:ID;onDelete:CASCADE"`
}

// 违章
type Violation struct {
	ID            uint    `gorm:"primaryKey"`
	Driver        Driver  `gorm:"foreignKey:ID;onDelete:CASCADE"`
	Vehicle       Vehicle `gorm:"foreignKey:ID;onDelete:CASCADE"`
	Team          Team    `gorm:"foreignKey:ID;onDelete:CASCADE"`
	Route         Route   `gorm:"foreignKey:ID;onDelete:CASCADE"`
	OccurredAt    time.Time
	ViolationType string
}

// 车辆
type Vehicle struct {
	ID  uint `gorm:"primaryKey"`
	VIN string
}

var db *gorm.DB

func init_db() error {
	dbConfig := viper.GetStringMap("database")
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig["user"], dbConfig["password"],
		dbConfig["host"], dbConfig["port"], dbConfig["name"])
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	if migrate := dbConfig["migrate"]; migrate == false {
		return nil
	}
	if err = db.AutoMigrate(&RoadManager{}, &Violation{}, &Driver{}, &Route{}, &Team{}, &Company{}, &Vehicle{}); err != nil {
		return err
	}
	return nil
}
