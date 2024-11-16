package main

import (
	"github.com/alihaqberdi/goga_go/internal/config"
	"github.com/alihaqberdi/goga_go/internal/handler"
	"github.com/alihaqberdi/goga_go/internal/pkg/postgres"
	"github.com/alihaqberdi/goga_go/internal/repo"
	"github.com/alihaqberdi/goga_go/internal/service"
	"github.com/alihaqberdi/goga_go/internal/service/caching"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand/v2"
	"net/http"
	"time"
)

func main() {
	db, err := postgres.ConnectDB(config.POSTGRES_URI)
	if err != nil {
		panic(err)
	}

	err = postgres.AutoMigrate(db)
	if err != nil {
		panic(err)
	}

	repos := repo.New(db)
	cache := caching.New()
	services := service.New(repos, cache)
	handlers := handler.New(services, cache)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	mw := handlers.MW
	_ = mw
	// api
	{
		r.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message":    "Server is running!!!",
				"random_num": rand.Int() % 1_000,
			})
		})

		v1 := r.Group("/api/v1")

		auth := v1.Group("/auth")
		{
			h := handlers.Auth

			auth.POST("/register", h.Register)
			auth.POST("/login", h.Login)
		}

	}

	log.Fatalln(r.Run(":" + config.PORT))

}
