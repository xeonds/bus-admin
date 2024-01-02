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
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Name      string
	CreatedAt time.Time
}

// 车队
type Team struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Name        string
	CompanyID   uint
	Company     Company `gorm:"foreignKey:CompanyID;onDelete:CASCADE"`
	ManagerName string
}

// 路线
type Route struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	Name   string
	TeamID uint
	Team   Team `gorm:"foreignKey:TeamID;onDelete:CASCADE"`
}

// 司机
type Driver struct {
	ID      uint `gorm:"primaryKey;autoIncrement"`
	Name    string
	RouteID uint
	Route   Route `gorm:"foreignKey:RouteID;onDelete:CASCADE"`
}

// 队长
type RoadManager struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	RouteID uint
	Route   Route `gorm:"foreignKey:RouteID;onDelete:CASCADE"`
}

// 违章
type Violation struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	DriverID      uint
	VehicleID     uint
	TeamID        uint
	RouteID       uint
	OccurredAt    time.Time
	ViolationType string

	Driver  Driver  `gorm:"foreignKey:DriverID;onDelete:CASCADE"`
	Vehicle Vehicle `gorm:"foreignKey:VehicleID;onDelete:CASCADE"`
	Team    Team    `gorm:"foreignKey:TeamID;onDelete:CASCADE"`
	Route   Route   `gorm:"foreignKey:RouteID;onDelete:CASCADE"`
}

// 车辆
type Vehicle struct {
	ID  uint `gorm:"primaryKey;autoIncrement"`
	VIN string
}

var db *gorm.DB

func init_db() error {
	dbConfig := viper.GetStringMap("database")
	var err error
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig["user"], dbConfig["password"],
		dbConfig["host"], dbConfig["port"], dbConfig["name"])
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
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
