package game

import (
	"github.com/absentbird/RPG-Consent-Checklist/internal/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateGame(c *gin.Context) {
    cl := Checklist{}
	c.JSON(http.StatusOK, cl)
}

func JoinGame(c *gin.Context) {
    cl := Checklist{}
	c.JSON(http.StatusOK, cl)
}

func UpdatePlayer(c *gin.Context) {
    cl := Checklist{}
	c.JSON(http.StatusOK, cl)
}
