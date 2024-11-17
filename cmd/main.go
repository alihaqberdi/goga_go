package main

import (
	"log"
	"math/rand/v2"
	"net/http"
	"time"

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
	handlers := handler.New(services, cache, jwtManager)

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
	mwClient := handlers.MW.AuthByRoles(types.UserRoleClient)
	mwContractor := handlers.MW.AuthByRoles(types.UserRoleContractor)
	_, _ = mwContractor, mwClient
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
		bids := r.Group("/api")
		{
			h := handlers.Bids
			bids.POST("/contractor/tenders/:tender_id/bid", h.CreateBid)
			bids.GET("/client/tenders/:tender_id/bids", h.GetList)
			bids.POST("/client/tenders/:tender_id/award:id", h.AwardBid)
		}

	}

	log.Fatalln(r.Run(":" + config.PORT))

}
