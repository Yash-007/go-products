package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	samplebackend "github.com/sample/temp/api/sample-backend"
)

func ReqHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("working fine")
}

func GinHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "everything working fine",
	})
}

func BuildServer(Addr string) *http.Server {
	router := gin.New()
	router.GET("/", GinHandler)
	newRouter := router.Group("/")

	samplebackend.RouteApis(newRouter)

	server := &http.Server{
		Addr:           Addr,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return server
	// router.Run(":8001")
}
