package handlers

import (
	"log"
	"net/http"
	"strconv"
	"twiteer/internal/data/postgres/repository"
	"twiteer/internal/domain/models"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	repo *repository.CommRepoSQL
}

func NewCommentHandler(repo *repository.CommRepoSQL) *CommentHandler {
	return &CommentHandler{repo: repo}
}

// POST /comments
func (h *CommentHandler) CreateComment(c *gin.Context) {
	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		log.Println("[ERROR] Invalid input:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if err := h.repo.Create(comment); err != nil {
		log.Println("[ERROR] Failed to create comment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create comment"})
		return
	}

	log.Printf("[SUCCESS] Comment created by user %d on tret %d\n", comment.UserID, comment.TretID)
	c.JSON(http.StatusCreated, gin.H{"message": "comment created"})
}

// PUT /comments/:id
func (h *CommentHandler) UpdateComment(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		log.Println("[ERROR] Invalid comment ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid comment id"})
		return
	}

	var comment models.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		log.Println("[ERROR] Invalid input:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	comment.ID = uint(id)
	if err := h.repo.Update(comment); err != nil {
		log.Println("[ERROR] Failed to update comment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update comment"})
		return
	}

	log.Printf("[SUCCESS] Comment %d updated\n", comment.ID)
	c.JSON(http.StatusOK, gin.H{"message": "comment updated"})
}

// DELETE /comments/:id
func (h *CommentHandler) DeleteComment(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.Delete(id); err != nil {
		log.Println("[ERROR] Failed to delete comment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete comment"})
		return
	}

	log.Printf("[SUCCESS] Comment %s deleted\n", id)
	c.JSON(http.StatusOK, gin.H{"message": "comment deleted"})
}

// GET /comments/:id
func (h *CommentHandler) GetComment(c *gin.Context) {
	id := c.Param("id")
	comment, err := h.repo.GetComment(id)
	if err != nil {
		log.Println("[ERROR] Failed to get comment:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "comment not found"})
		return
	}

	log.Printf("[SUCCESS] Comment %s retrieved\n", id)
	c.JSON(http.StatusOK, comment)
}

// PATCH /comments/:id/like
func (h *CommentHandler) LikeComment(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.LikeComment(id); err != nil {
		log.Println("[ERROR] Failed to like comment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to like comment"})
		return
	}

	log.Printf("[SUCCESS] Comment %s liked\n", id)
	c.JSON(http.StatusOK, gin.H{"message": "comment liked"})
}

// PATCH /comments/:id/dislike
func (h *CommentHandler) DislikeComment(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DislikeComment(id); err != nil {
		log.Println("[ERROR] Failed to dislike comment:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to dislike comment"})
		return
	}

	log.Printf("[SUCCESS] Comment %s disliked\n", id)
	c.JSON(http.StatusOK, gin.H{"message": "comment disliked"})
}

// GET /trets/:id/comments
func (h *CommentHandler) GetCommentsByTret(c *gin.Context) {
	tretID := c.Param("id")
	comments, err := h.repo.GetCommentsByTretID(tretID)
	if err != nil {
		log.Printf("[ERROR] Cannot get comments for tret %s: %v", tretID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.Printf("[SUCCESS] Retrieved %d comments for tret %s", len(comments), tretID)
	c.JSON(http.StatusOK, comments)
}
