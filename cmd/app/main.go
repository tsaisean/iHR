package main

import (
	"github.com/gin-gonic/gin"
	"iHR/config"
	"iHR/db"
	"iHR/redis"
	"iHR/route"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	db.Connect(&cfg.Database)

	redis.Connect(cfg.Redis)

	// AutoMigrations
	db.AutoMigrate(db.DB)

	r := gin.Default()

	route.RegisterRoutes(r, cfg)

	r.Run()
}
