package server

import (
	"log"
	"net/url"

	"github.com/charmbracelet/charm/kv"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func Start() error {
	db, err := kv.OpenWithDefaults("grove")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Sync(); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	router.Use(gin.Recovery())
	// router.SetTrustedProxies([]string{"0.0.0.0"})
	//Routes
	router.GET("/plants/:package", func(ctx *gin.Context) {
		pk := url.QueryEscape(ctx.Param("package"))
		data, err := db.Get([]byte(pk))
		if err != nil {
			ctx.AbortWithError(503, err)
			return
		}
		ctx.Data(200, "application/octet-stream", data)
	})

	router.Run(":8080")

	return nil
}
