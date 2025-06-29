package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"twiteer/internal/data/postgres/repository"
)

type ReactionHandler struct {
	reactionRepo *repository.ReactionRepoSQL
	tretRepo     *repository.TretRepoSQL
}

func NewReactionHandler(reactionRepo *repository.ReactionRepoSQL, tretRepo *repository.TretRepoSQL) *ReactionHandler {
	return &ReactionHandler{
		reactionRepo: reactionRepo,
		tretRepo:     tretRepo,
	}
}

// PATCH /trets/:id/like/:userId
func (h *ReactionHandler) LikeTret(c *gin.Context) {
	tretID := c.Param("id")
	userID := c.Param("userId")
	log.Printf("[LIKE TRET] UserID: %s, TretID: %s", userID, tretID)

	if err := h.reactionRepo.LikeTret(userID, tretID); err != nil {
		log.Printf("[ERROR] Failed to like tret (reaction): %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := h.tretRepo.LikeTret(tretID); err != nil {
		log.Printf("[ERROR] Failed to like tret (increment): %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("[SUCCESS] Tret liked by user %s", userID)
	c.JSON(http.StatusOK, gin.H{"message": "Tret liked"})
}

// PATCH /trets/:id/dislike/:userId
func (h *ReactionHandler) DislikeTret(c *gin.Context) {
	tretID := c.Param("id")
	userID := c.Param("userId")
	log.Printf("[DISLIKE TRET] UserID: %s, TretID: %s", userID, tretID)

	if err := h.reactionRepo.DislikeTret(userID, tretID); err != nil {
		log.Printf("[ERROR] Failed to dislike tret (reaction): %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := h.tretRepo.DislikeTret(tretID); err != nil {
		log.Printf("[ERROR] Failed to dislike tret (increment): %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("[SUCCESS] Tret disliked by user %s", userID)
	c.JSON(http.StatusOK, gin.H{"message": "Tret disliked"})
}

// GET /trets/:id/liked-users
func (h *ReactionHandler) GetLikedUsers(c *gin.Context) {
	tretID := c.Param("id")
	log.Printf("[GET LIKED USERS] TretID: %s", tretID)

	users, err := h.reactionRepo.GetLikedUsers(tretID)
	if err != nil {
		log.Printf("[ERROR] Failed to fetch liked users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("[SUCCESS] Found %d users who liked tret %s", len(users), tretID)
	c.JSON(http.StatusOK, users)
}

// GET /trets/:id/disliked-users
func (h *ReactionHandler) GetDislikedUsers(c *gin.Context) {
	tretID := c.Param("id")
	log.Printf("[GET DISLIKED USERS] TretID: %s", tretID)

	users, err := h.reactionRepo.GetDislikedUsers(tretID)
	if err != nil {
		log.Printf("[ERROR] Failed to fetch disliked users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("[SUCCESS] Found %d users who disliked tret %s", len(users), tretID)
	c.JSON(http.StatusOK, users)
}

// GET /users/:id/liked-trets
func (h *ReactionHandler) GetLikedTrets(c *gin.Context) {
	userID := c.Param("id")
	log.Printf("[GET LIKED TRETS] UserID: %s", userID)

	trets, err := h.reactionRepo.GetLikedTrets(userID)
	if err != nil {
		log.Printf("[ERROR] Failed to fetch liked trets: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("[SUCCESS] Found %d trets liked by user %s", len(trets), userID)
	c.JSON(http.StatusOK, trets)
}

// GET /users/:id/disliked-trets
func (h *ReactionHandler) GetDislikedTrets(c *gin.Context) {
	userID := c.Param("id")
	log.Printf("[GET DISLIKED TRETS] UserID: %s", userID)

	trets, err := h.reactionRepo.GetDislikedTrets(userID)
	if err != nil {
		log.Printf("[ERROR] Failed to fetch disliked trets: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("[SUCCESS] Found %d trets disliked by user %s", len(trets), userID)
	c.JSON(http.StatusOK, trets)
}
