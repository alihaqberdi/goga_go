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

		// Tenders
		client := r.Group("/api/client", mwClient)
		{
			h := handlers.Tender

			client.POST("/tenders", h.CreateTender)
			client.GET("/tenders", h.GetListTenders)
			client.PUT("/tenders", h.UpdateTender)
			client.GET("/tenders/:client_id", h.GetListTendersByUser)
		}

		// Bids
		bids_client := r.Group("/api/client")
		bids_contractor := r.Group("/api/contractor")
		{
			h := handlers.Bids
			bids_contractor.POST("/tenders/:tender_id/bid", h.Create)
			bids_contractor.DELETE("/tenders/:tender_id/bid/:id", h.Delete)
			bids_client.GET("/tenders/:tender_id/bids", h.GetList)
			bids_client.POST("/tenders/:tender_id/award:id", h.AwardBid)
		}
		users := r.Group("/users", mwClient)
		{
			h := handlers.Bids
			users.GET("/:id/bids", h.UserBids)
		}

	}

	log.Fatalln(r.Run(":" + config.PORT))

}
