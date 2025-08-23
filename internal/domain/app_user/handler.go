package appuser

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marriosdev/export-api/internal/auth"
	"golang.org/x/crypto/bcrypt"
)

type Handler struct{}

var fakeUser = AppUser{
	ID:       "Asad",
	Username: "admin",
	// senha "123456"
	Password: "$2a$10$0H00C5oGk8qRupW5lR8zjO2v8XqT5L4lE03bW3xG1mR2zG0lI2o1u",
}

func loginHandler(c *gin.Context) {
	var dto loginDTO
	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "payload inválido"})
		return
	}

	if dto.Username != fakeUser.Username || bcrypt.CompareHashAndPassword([]byte(fakeUser.Password), []byte(dto.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "credenciais inválidas"})
		return
	}

	token, err := auth.GenerateToken(auth.UserAuthPayload{
		ID:       fakeUser.ID,
		Username: fakeUser.Username,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "não foi possível gerar token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": token})
}

func meHandler(c *gin.Context) {
	claims := c.MustGet("claims").(*auth.CustomClaims)
	c.JSON(http.StatusOK, gin.H{
		"user_id":  claims.UserID,
		"username": claims.Username,
		"exp":      claims.ExpiresAt.Time,
	})
}
