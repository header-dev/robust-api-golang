package auth

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

func AccessToken(c *gin.Context) {

	expr := jwt.NewTime(float64(time.Now().Add(5 * time.Minute).Unix()))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: expr,
	})

	ss, err := token.SignedString([]byte("==signature=="))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": ss,
	})
}
