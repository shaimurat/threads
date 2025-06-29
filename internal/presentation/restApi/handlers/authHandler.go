package handlers

import (
	"log"
	"net/http"
	"twiteer/config"
	"twiteer/internal/data/postgres/repository"
	"twiteer/utils"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userRepo *repository.UserRepoSQL
}

func NewAuthHandler(userRepo *repository.UserRepoSQL) *AuthHandler {
	return &AuthHandler{userRepo: userRepo}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := h.userRepo.GetUserByEmail(input.Email)
	if err != nil {
		log.Printf("User not found: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if !utils.CheckPasswordHash(input.Password, user.Password) {
		log.Println("Invalid password")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	session, err := config.Store.Get(c.Request, "users")
	if err != nil {
		log.Printf("Session error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create session"})
		return
	}

	session.Values["userID"] = user.ID
	if err := session.Save(c.Request, c.Writer); err != nil {
		log.Printf("Failed to save session: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	session, err := config.GetSession(c.Request)
	if err != nil {
		log.Printf("[LOGOUT ERROR] Failed to get session: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not log out"})
		return
	}

	session.Options.MaxAge = -1 // Delete the session
	if err := session.Save(c.Request, c.Writer); err != nil {
		log.Printf("[LOGOUT ERROR] Failed to save session: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not log out"})
		return
	}

	log.Println("[LOGOUT SUCCESS] User logged out successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}
