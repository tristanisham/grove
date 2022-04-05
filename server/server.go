package server

import (
	"log"

	"github.com/charmbracelet/charm/kv"
	"github.com/gin-gonic/gin"
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
	

	router.Run(":8080")

	return nil
}
