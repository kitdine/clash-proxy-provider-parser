package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kitdine/clash-proxy-provider-parser/clash"
	"log"
	"net/http"
)

// the entrance of a web service
func Serv(port string) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "PONG")
	})
	router.GET("/clash/parse", clash.Parse)

	err := router.Run(port)
	if err != nil {
		log.Println(err.Error())
	}
}
