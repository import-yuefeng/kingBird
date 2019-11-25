package kinglet

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	PORT = ":2048"
)

func init() {
}

func Kinglet(configPath string) {
	log.Info("hello, now is kinglet mode")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(PORT) // listen and serve on 0.0.0.0:2048

}
