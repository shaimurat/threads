package repository

import (
	"log"
	"strconv"

	"gorm.io/gorm"
	"twiteer/internal/data/postgres/mapper"
	postgreModels "twiteer/internal/data/postgres/models"
	"twiteer/internal/domain/models"
	"twiteer/utils"
)

type UserRepoSQL struct {
	Db *gorm.DB
}

func NewUserRepoSQL(db *gorm.DB) *UserRepoSQL {
	return &UserRepoSQL{
		Db: db,
	}
}
func (r *UserRepoSQL) Create(user models.User) error {
	log.Printf("[UserRepoSQL.Create] Creating user: %s", user.Username)

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Printf("[UserRepoSQL.Create] Error hashing password: %v", err)
		return err
	}
	user.Password = string(hashedPassword)

	doc := mapper.FromUser(user)
	result := r.Db.Create(&doc)

	if result.Error != nil {
		log.Printf("[UserRepoSQL.Create] Error saving to DB: %v", result.Error)
		return result.Error
	}

	log.Printf("[UserRepoSQL.Create] User created with ID: %d", doc.ID)
	return nil
}

func (r *UserRepoSQL) Update(user models.User) error {
	log.Printf("[UserRepoSQL.Update] Updating user ID: %d", user.ID)

	doc := mapper.FromUser(user)
	result := r.Db.Save(&doc)

	if result.Error != nil {
		log.Printf("[UserRepoSQL.Update] Error updating DB: %v", result.Error)
		return result.Error
	}

	log.Printf("[UserRepoSQL.Update] User updated successfully")
	return nil
}

func (r *UserRepoSQL) Delete(id string) error {
	log.Printf("[UserRepoSQL.Delete] Deleting user ID: %s", id)

	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Printf("[UserRepoSQL.Delete] Error parsing ID: %v", err)
		return err
	}

	result := r.Db.Delete(&postgreModels.UserDoc{}, uid)
	if result.Error != nil {
		log.Printf("[UserRepoSQL.Delete] Error deleting user: %v", result.Error)
		return result.Error
	}

	log.Printf("[UserRepoSQL.Delete] User deleted successfully")
	return nil
}

func (r *UserRepoSQL) GetUser(id string) (models.User, error) {
	log.Printf("[UserRepoSQL.GetUser] Fetching user ID: %s", id)

	uid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Printf("[UserRepoSQL.GetUser] Error parsing ID: %v", err)
		return models.User{}, err
	}

	var doc postgreModels.UserDoc
	result := r.Db.First(&doc, uid)
	if result.Error != nil {
		log.Printf("[UserRepoSQL.GetUser] Error retrieving user: %v", result.Error)
		return models.User{}, result.Error
	}

	log.Printf("[UserRepoSQL.GetUser] User fetched: %s", doc.Username)
	return mapper.ToUser(doc), nil
}
func (r *UserRepoSQL) GetUserByEmail(email string) (models.User, error) {
	var doc postgreModels.UserDoc
	result := r.Db.Where("email = ?", email).First(&doc)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return mapper.ToUser(doc), nil
}
