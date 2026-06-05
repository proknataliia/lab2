package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"ipinfo-http-server/internal/config"
	"ipinfo-http-server/internal/httpapi"
	"ipinfo-http-server/internal/ipinfo"
	"ipinfo-http-server/internal/ratelimit"
)

func main() {
	cfg := config.Load()

	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Static("/client", "./web")

	server := httpapi.NewServer(
		ipinfo.NewClient(cfg.IPInfoAPIKey),
		ratelimit.New(cfg.RateLimitSec),
	)
	server.RegisterRoutes(r)

	log.Printf("Server started on http://localhost:%s", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
