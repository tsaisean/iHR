package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"iHR/config"
	"iHR/repositories/db"
	"iHR/repositories/redis"
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
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("iHR-session", store))

	route.RegisterRoutes(r, cfg)

	r.Run()
}
