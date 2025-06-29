package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"twiteer/internal/data/postgres/repository"
	"twiteer/internal/domain/models"
)

type TretHandler struct {
	repo *repository.TretRepoSQL
}

func NewTretHandler(repo *repository.TretRepoSQL) *TretHandler {
	return &TretHandler{repo: repo}
}

// POST /trets
func (h *TretHandler) CreateTret(c *gin.Context) {
	var tret models.Tret
	if err := c.ShouldBindJSON(&tret); err != nil {
		log.Printf("[CreateTret] Invalid input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	log.Printf("[CreateTret] Creating tret: %+v", tret)

	if err := h.repo.Create(tret); err != nil {
		log.Printf("[CreateTret] Error creating tret: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("[CreateTret] Successfully created tret")
	c.JSON(http.StatusCreated, gin.H{"message": "Tret created"})
}

// PUT /trets/:id
func (h *TretHandler) UpdateTret(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Printf("[UpdateTret] Invalid ID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var tret models.Tret
	if err := c.ShouldBindJSON(&tret); err != nil {
		log.Printf("[UpdateTret] Invalid input: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	tret.ID = uint(id)
	log.Printf("[UpdateTret] Updating tret ID %d: %+v", id, tret)

	if err := h.repo.Update(tret); err != nil {
		log.Printf("[UpdateTret] Error updating tret: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("[UpdateTret] Successfully updated tret")
	c.JSON(http.StatusOK, gin.H{"message": "Tret updated"})
}

// DELETE /trets/:id
func (h *TretHandler) DeleteTret(c *gin.Context) {
	id := c.Param("id")
	log.Printf("[DeleteTret] Deleting tret ID: %s", id)

	if err := h.repo.Delete(id); err != nil {
		log.Printf("[DeleteTret] Error deleting tret: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("[DeleteTret] Successfully deleted tret")
	c.JSON(http.StatusOK, gin.H{"message": "Tret deleted"})
}

// GET /trets/:id
func (h *TretHandler) GetTret(c *gin.Context) {
	id := c.Param("id")
	log.Printf("[GetTret] Fetching tret ID: %s", id)

	tret, err := h.repo.GetTret(id)
	if err != nil {
		log.Printf("[GetTret] Error fetching tret: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "tret not found"})
		return
	}

	log.Printf("[GetTret] Fetched tret: %+v", tret)
	c.JSON(http.StatusOK, tret)
}

// GET /users/:id/trets
func (h *TretHandler) GetUserTrets(c *gin.Context) {
	userID := c.Param("id")
	log.Printf("[GetUserTrets] Fetching trets for user ID: %s", userID)

	trets, err := h.repo.GetUserTrets(userID)
	if err != nil {
		log.Printf("[GetUserTrets] Error fetching user trets: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("[GetUserTrets] Found %d trets", len(trets))
	c.JSON(http.StatusOK, trets)
}

// PATCH /trets/:id/like
func (h *TretHandler) LikeTret(c *gin.Context) {
	id := c.Param("id")
	log.Printf("[LikeTret] Liking tret ID: %s", id)

	if err := h.repo.LikeTret(id); err != nil {
		log.Printf("[LikeTret] Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("[LikeTret] Like successful")
	c.JSON(http.StatusOK, gin.H{"message": "liked"})
}

// PATCH /trets/:id/dislike
func (h *TretHandler) DislikeTret(c *gin.Context) {
	id := c.Param("id")
	log.Printf("[DislikeTret] Disliking tret ID: %s", id)

	if err := h.repo.DislikeTret(id); err != nil {
		log.Printf("[DislikeTret] Error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("[DislikeTret] Dislike successful")
	c.JSON(http.StatusOK, gin.H{"message": "disliked"})
}
