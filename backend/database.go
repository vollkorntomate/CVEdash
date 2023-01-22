package main

import (
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var DB Database

type Database struct {
	DB *gorm.DB
}

func NewDB(path string) Database {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatalln("Could not open database connection to", path)
	}
	return Database{DB: db}
}

func (db *Database) InitDB() {
	db.DB.AutoMigrate(&MinimalCVEData{})
	db.DB.AutoMigrate(&LastNVDUpdate{})
}

func (db *Database) GetLastNVDUpdate() time.Time {
	lastUpdate := LastNVDUpdate{}
	var numEntries int64

	db.DB.Model(&lastUpdate).Count(&numEntries)

	if numEntries > 0 {
		db.DB.First(&lastUpdate)
	}

	return lastUpdate.LastUpdate
}

func (db *Database) UpdateLastNVDUpdate() {
	lastUpdate := LastNVDUpdate{LastUpdate: time.Now()}

	db.DB.Where("1=1").Save(&lastUpdate) // override existing single value
}

func (db *Database) SaveBatch(data []MinimalCVEData) {
	db.DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(data)
}
