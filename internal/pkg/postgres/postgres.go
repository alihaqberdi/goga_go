package postgres

import (
	"github.com/alihaqberdi/goga_go/internal/config"
	"github.com/alihaqberdi/goga_go/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {

	if config.POSTGRES_DROP_TABELS {
		db.Migrator().DropTable(&models.User{})
		db.Migrator().DropTable(&models.User{})
		db.Migrator().DropTable(&models.Tender{})
		db.Migrator().DropTable(&models.Bid{})
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Tender{})
	db.AutoMigrate(&models.Bid{})

	return nil
}
