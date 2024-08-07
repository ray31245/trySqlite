package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TempArticle struct {
	Title   string
	Content string
}

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error while connecting to database: %v", err)
	}

	err = db.AutoMigrate(&TempArticle{})
	if err != nil {
		log.Fatalf("Error while migrating: %v", err)
	}
	err = db.Create(&TempArticle{Title: "Hello", Content: "World"}).Error
	if err != nil {
		log.Fatalf("Error while creating: %v", err)
	}
	var article TempArticle
	err = db.First(&article).Error
	if err != nil {
		log.Fatalf("Error while fetching: %v", err)
	}
	log.Printf("Article: %v", article)
}
