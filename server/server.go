package server

import (
	"fmt"
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
		// data, err2 := db.Get([]byte(pk + "@" + version))
		tar_file := os.Getenv("GROVE_PKG_DIR") + pk + "/" + version + "/" + pk + "@" + version + ".tar.gz"
		log.Println(os.Getenv("GROVE_PKG_DIR"))
		if _, err := os.Stat(tar_file); os.IsNotExist(err) {
			ctx.AbortWithError(503, fmt.Errorf("package requested is not found on this server. %s", os.Getenv("GROVE_REPO")))
			return
		}
		// app.Tar(os.Getenv("GROVE_PKG_DIR") + pk + "/" + version, os.Getenv("GROVE_PKG_DIR") + pk + "/" + version+"/"+pk+"@"+version)

		ctx.File(tar_file)
	})

	router.Run(":" + os.Getenv("GROVE_SERVER_PORT"))

	return nil
}
