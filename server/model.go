package main

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Company struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	Contact   string
}

type Team struct {
	ID          uint `gorm:"primaryKey"`
	Name        string
	Company     Company `gorm:"foreignKey:ID;onDelete:CASCADE"`
	ManagerName string
}

type Route struct {
	ID   uint `gorm:"primaryKey"`
	Name string
	Team Team `gorm:"foreignKey:ID;onDelete:CASCADE"`
}

type Driver struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Route Route `gorm:"foreignKey:ID;onDelete:CASCADE"`
}

type RoadManager struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Route Route `gorm:"foreignKey:ID;onDelete:CASCADE"`
}

type Violation struct {
	ID            uint    `gorm:"primaryKey"`
	Driver        Driver  `gorm:"foreignKey:ID;onDelete:CASCADE"`
	Vehicle       Vehicle `gorm:"foreignKey:ID;onDelete:CASCADE"`
	Team          Team    `gorm:"foreignKey:ID;onDelete:CASCADE"`
	Route         Route   `gorm:"foreignKey:ID;onDelete:CASCADE"`
	Station       Station `gorm:"foreignKey:ID;onDelete:CASCADE"`
	OccurredAt    time.Time
	ViolationType string
}

type Vehicle struct {
	ID  uint `gorm:"primaryKey"`
	VIN string
}

type Station struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

var db *gorm.DB

func init_db() {
	// 连接到数据库
	db, err := gorm.Open(mysql.Open("username:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	// 创建表
	err = db.AutoMigrate(&Company{}, &Team{}, &Route{}, &Driver{}, &RoadManager{}, &Violation{}, &Vehicle{}, &Station{})
	if err != nil {
		panic("failed to create tables")
	}
}
