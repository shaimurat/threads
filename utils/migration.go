package utils

import (
	"gorm.io/gorm"
	"log"
	"twiteer/internal/data/postgres/models"
)

func AutoMigration(db *gorm.DB) error {
	err := db.AutoMigrate(&postgreModels.UserDoc{})
	log.Println("Migrating users table")
	if err != nil {
		log.Fatalf("Error when auto migrating UserTable:%v", err)
		return err
	}
	err = db.AutoMigrate(&postgreModels.TretDoc{})
	log.Println("Migrating trets table")
	if err != nil {
		log.Fatalf("Error when auto migrating TretTable:%v", err)
		return err
	}
	err = db.AutoMigrate(postgreModels.CommentDoc{})
	log.Println("Migrating comment table")
	if err != nil {
		log.Fatalf("Error when auto migrating CommentTable:%v", err)
		return err
	}
	err = db.AutoMigrate(postgreModels.ReactionDoc{})
	log.Println("Migrating reaction table")
	if err != nil {
		log.Fatalf("Error when auto migrating ReactionTable:%v", err)
		return err
	}
	return nil
}
