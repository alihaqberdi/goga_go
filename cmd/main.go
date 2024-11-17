package main

import (
	"log"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/alihaqberdi/goga_go/internal/handler/mw"

	"github.com/alihaqberdi/goga_go/internal/models/types"

	_ "github.com/alihaqberdi/goga_go/docs"

	"github.com/alihaqberdi/goga_go/internal/config"
	"github.com/alihaqberdi/goga_go/internal/handler"
	"github.com/alihaqberdi/goga_go/internal/pkg/jwt_manager"
	"github.com/alihaqberdi/goga_go/internal/pkg/postgres"
	"github.com/alihaqberdi/goga_go/internal/repo"
	"github.com/alihaqberdi/goga_go/internal/service"
	"github.com/alihaqberdi/goga_go/internal/service/caching"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go API Example
// @version 1.0
// @description API documentation for the Go application
// @host localhost:8080
// @BasePath /api/v1
func main() {

	db, err := postgres.ConnectDB(config.POSTGRES_URI)
	if err != nil {
		panic(err)
	}

	err = postgres.AutoMigrate(db)
	if err != nil {
		panic(err)
	}

	jwtManager := jwt_manager.New(config.JWT_SIGNING_KEY, config.JWT_EXPIRY_DURATION)
	repos := repo.New(db)
	cache := caching.New()
	services := service.New(repos, cache, jwtManager)
	mw := mw.New(services, cache, jwtManager)
	handlers := handler.New(services, cache, mw)

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
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	mwClient := mw.AuthByRoles(types.UserRoleClient)
	mwContractor := mw.AuthByRoles(types.UserRoleContractor)
	// api
	{
		r.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message":    "Server is running!!!",
				"random_num": rand.Int() % 1_000,
			})
		})

		// Auth
		{
			h := handlers.Auth

			r.POST("/register", h.Register)
			r.POST("/login", h.Login)
		}

		// Tenders
		client := r.Group("/api/client", mwClient)
		{
			h := handlers.Tenders

			client.POST("/tenders", h.Create)
			client.GET("/tenders", h.GetListByClient)
			client.PUT("/tenders/:id", h.Update)
			client.DELETE("/tenders/:id", h.Delete)
			client.GET("/tenders/:id", h.GetListTendersByUser)
		}

		// Bids
		contractor := r.Group("/api/contractor", mwContractor)
		{
			h := handlers.Bids
			contractor.POST("/tenders/:tender_id/bid", h.Create)
			contractor.GET("/bids", h.GetListByContractor)
			contractor.DELETE("/bids/:id", h.Delete)

			client.GET("/tenders/:id/bids", h.GetList)
			client.POST("/tenders/:id/award:id", h.AwardBid)
		}
	}

	log.Fatalln(r.Run(":" + config.PORT))

}
