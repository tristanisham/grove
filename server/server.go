package server

import (
	
	"log"
	"net/url"
	"os"

	"github.com/charmbracelet/charm/kv"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)
//Start begins the webserver in the 'server' module.
func Start(allowed_proxies []string) error {
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
	router.SetTrustedProxies(allowed_proxies)
	//Routes
	router.GET("/plants/:package/:version", func(ctx *gin.Context) {
		pk := url.QueryEscape(ctx.Param("package"))
		version := url.QueryEscape(ctx.Param("version"))
		log.Print(pk, "@", version)
		data, err := db.Get([]byte(pk + "@" + version))
		if err != nil {
			ctx.AbortWithError(503, err)
			return
		}
		ctx.Data(200, "application/octet-stream", data)
	})

	router.Run(":" + os.Getenv("GROVE_SERVER_PORT"))

	return nil
}
