package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"twiteer/config"
	"twiteer/internal/data/postgres/repository"
	"twiteer/internal/presentation/restApi/handlers"
	"twiteer/middleware"
	"twiteer/utils"
)

func main() {
	utils.LoadEnv()
	db := config.ConnectPostgres()
	err := utils.AutoMigration(db)
	if err != nil {
		log.Println(err)
	}
	r := gin.Default()
	userRepo := &repository.UserRepoSQL{Db: db}
	tretRepo := repository.NewTretRepoSQL(db)
	reactionRepo := repository.NewReactionRepoSQL(db)
	commentRepo := repository.NewCommRepoSQL(db)

	userHandler := handlers.NewUserHandler(userRepo)
	tretHandler := handlers.NewTretHandler(tretRepo)
	reactionHandler := handlers.NewReactionHandler(reactionRepo, tretRepo)
	commentHandler := handlers.NewCommentHandler(commentRepo)
	authHandler := handlers.NewAuthHandler(userRepo)

	// PUBLIC ROUTES
	r.POST("/login", authHandler.Login)

	r.GET("/users/:id", userHandler.GetUser)
	r.GET("/trets/:id", tretHandler.GetTret)
	r.GET("/users/:id/trets", tretHandler.GetUserTrets)
	r.GET("/trets/:id/comments", commentHandler.GetCommentsByTret)
	r.GET("/users/:id/liked-trets", reactionHandler.GetLikedTrets)
	r.GET("/users/:id/disliked-trets", reactionHandler.GetDislikedTrets)
	r.GET("/trets/:id/liked-users", reactionHandler.GetLikedUsers)
	r.GET("/trets/:id/disliked-users", reactionHandler.GetDislikedUsers)
	r.GET("/comments/:id", commentHandler.GetComment)
	r.POST("/users", userHandler.CreateUser)
	// PROTECTED ROUTES
	auth := r.Group("/")
	auth.Use(middleware.AuthRequired())
	{
		auth.POST("/logout", authHandler.Logout)

		// USERS
		auth.PUT("/users/:id", userHandler.UpdateUser)
		auth.DELETE("/users/:id", userHandler.DeleteUser)

		// TRETS
		auth.POST("/trets", tretHandler.CreateTret)
		auth.PUT("/trets/:id", tretHandler.UpdateTret)
		auth.DELETE("/trets/:id", tretHandler.DeleteTret)

		// REACTIONS
		auth.PATCH("/trets/:id/like/:userId", reactionHandler.LikeTret)
		auth.PATCH("/trets/:id/dislike/:userId", reactionHandler.DislikeTret)

		// COMMENTS
		auth.POST("/comments", commentHandler.CreateComment)
		auth.PUT("/comments/:id", commentHandler.UpdateComment)
		auth.DELETE("/comments/:id", commentHandler.DeleteComment)
		auth.PATCH("/comments/:id/like", commentHandler.LikeComment)
		auth.PATCH("/comments/:id/dislike", commentHandler.DislikeComment)
	}

	r.Run(":8080") // localhost:8080
}
