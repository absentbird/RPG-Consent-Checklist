package main

import (
	"github.com/absentbird/RPG-Consent-Checklist/internal/game"
	"github.com/absentbird/RPG-Consent-Checklist/internal/util"
	"github.com/gin-gonic/gin"
)

const (
	PRODCONFIG = "configs/config.yaml"
	DEVCONFIG  = "configs/config-dev.yaml"
)

func main() {
	// Initialize the config, database, and router
	conf := util.Config{}
	conf.Load(PRODCONFIG, DEVCONFIG)
	util.DB = db.ConnectDB(conf.DBConn)
    if conf.Mode == "production" {
        gin.SetMode(gin.ReleaseMode)
    }
	r := gin.Default()

	r.POST("/game/new", game.CreateGame)
	r.POST("/game/:id/join", game.JoinGame)
    r.POST("/game/:id/update/:pid", game.UpdatePlayer)

	err := r.Run("127.0.0.1:" + conf.Port)
	util.Check(err, "Error starting API")
}
