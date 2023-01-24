package data

import (
	"log"
	"time"

	"vollkorntomate.de/cvedash-api/internal/tools"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const DB_VERSION = 1

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

func (db *Database) GetLastNVDUpdate() LastNVDUpdate {
	lastUpdate := LastNVDUpdate{}
	var numEntries int64

	db.DB.Model(&lastUpdate).Count(&numEntries)

	if numEntries > 0 {
		db.DB.First(&lastUpdate)
	}

	return lastUpdate
}

func (db *Database) UpdateLastNVDUpdate(t time.Time) {
	lastUpdate := LastNVDUpdate{LastUpdate: t, Version: DB_VERSION}

	db.DB.Where("1=1").Save(&lastUpdate) // override existing single value
}

func (db *Database) SaveBatch(data []MinimalCVEData) {
	if len(data) > 0 {
		db.DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(data)
	}
}

func (db *Database) GetStats(since time.Time) CVEStats {
	query := "SELECT cvss_base_severity, count(*) as count FROM minimal_cve_data WHERE published > ? GROUP BY cvss_base_severity;"
	var stats CVEStats

	dbResults := make([]struct {
		CVSSBaseSeverity string `gorm:"cvss_base_severity"`
		Count            int    `gorm:"count"`
	}, 0)

	db.DB.Raw(query, tools.FormatISODate(since)).Scan(&dbResults)

	for i := 0; i < len(dbResults); i++ {
		count := dbResults[i].Count
		switch dbResults[i].CVSSBaseSeverity {
		case "":
			stats.NumUnknownSeverity = count
		case "CRITICAL":
			stats.NumCritical = count
		case "HIGH":
			stats.NumHigh = count
		case "MEDIUM":
			stats.NumMedium = count
		case "LOW":
			stats.NumLow = count
		}
	}

	return stats
}
