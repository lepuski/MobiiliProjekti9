package main

import (
	"errors"
	"mobiiliprojekti/internal/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (app *application) registerPost(c *gin.Context) {
	var req loginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	err := app.users.Insert(req.Username, req.Password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateUsername) {
			c.JSON(409, gin.H{"error": "Username already exists"})
		} else {
			app.logger.Error(err.Error())
			c.JSON(500, gin.H{"error": "Could not register user"})
		}
		return
	}

	c.JSON(201, gin.H{"message": "User registered successfully"})
}

func (app *application) loginPost(c *gin.Context) {
	var req loginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	userID, err := app.users.Authenticate(req.Username, req.Password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			c.JSON(401, gin.H{"error": "Invalid credentials"})
		} else {
			app.logger.Error(err.Error())
			c.JSON(500, gin.H{"error": "Authentication failed"})
		}
		return
	}

	session := sessions.Default(c)
	session.Set("userID", userID)
	if err := session.Save(); err != nil {
		c.JSON(500, gin.H{"error": "Could not save session"})
		return
	}

	c.JSON(200, gin.H{"message": "Login successful"})
}

func (app *application) loginGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Send a POST request to /login with 'username' and 'password' to log in.",
	})
}

func (app *application) registerGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Send a POST request to /register with 'username', 'password', and 'profile_id' to register.",
	})
}

func (app *application) logout(c *gin.Context) {

	session := sessions.Default(c)
	session.Clear()

	if err := session.Save(); err != nil {
		c.JSON(500, gin.H{"error": "Could not clear session"})
		return
	}

	c.JSON(200, gin.H{"message": "Logout successful"})
}


func (app *application) addFavoriteTeam(c *gin.Context) {
	var req struct {
		Team string `json:"team" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	err := app.users.UpdateFavoriteTeam(userID.(int), req.Team)
	if err != nil {
		app.logger.Error(err.Error())
		c.JSON(500, gin.H{"error": "Could not update favorite team"})
		return
	}

	c.JSON(200, gin.H{"message": "Favorite team added successfully"})
}


func (app *application) removeFavoriteTeam(c *gin.Context) {
	session := sessions.Default(c)
	userID := session.Get("userID")
	if userID == nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		return
	}

	err := app.users.RemoveFavoriteTeam(userID.(int))
	if err != nil {
		app.logger.Error(err.Error())
		c.JSON(500, gin.H{"error": "Could not remove favorite team"})
		return
	}

	c.JSON(200, gin.H{"message": "Favorite team removed successfully"})
}


func (app *application) getFavoriteTeam(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		app.logger.Error("User ID not found in context")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	userID, ok := userIDInterface.(int)
	if !ok {
		app.logger.Error("User ID has incorrect type")
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	favTeam, err := app.users.GetFavoriteTeam(userID)
	if err != nil {
		if errors.Is(err, models.ErrUserNotFound) {
			c.JSON(404, gin.H{"error": "User not found"})
		} else {
			app.logger.Error("Error fetching favorite team:", err.Error())
			c.JSON(500, gin.H{"error": "Could not retrieve favorite team"})
		}
		return
	}

	if favTeam == "" {
		c.JSON(200, gin.H{"favorite_team": nil})
	} else {
		c.JSON(200, gin.H{"favorite_team": favTeam})
	}
}
