package user

import (
	"net/http"
	"register/jwt-api/orm"

	"github.com/gin-gonic/gin"
)

func ReadAll(c *gin.Context) {
	var users []orm.User
	orm.Db.Find(&users)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "users": users})
}

func Profile(c *gin.Context) {
	userID := c.MustGet("userID").(float64)
	var user orm.User
	orm.Db.First(&user,userID)
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Read Success", "users": user})
}
