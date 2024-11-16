package repo

import (
	"github.com/alihaqberdi/goga_go/internal/models"
	"gorm.io/gorm"
)

type Users struct {
	DB *gorm.DB
}

func (r *Users) Create(m *models.User) error {
	err := r.DB.Create(m).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Users) GetByID(id uint) (*models.User, error) {
	var m models.User

	err := r.DB.First(&m, id).Error
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *Users) GetList(limit, offset int) ([]models.User, error) {
	var list []models.User
	query := r.DB.Limit(limit).Offset(offset)

	err := query.Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

func (r *Users) Update(tender *models.User) error {
	return r.DB.Save(tender).Error
}

func (r *Users) Delete(id uint) error {
	return r.DB.Delete(&models.User{}, id).Error
}

//func (r *Users) GetByEmailOrUsername(email string, username string) (*models.User, error) {
//	var m models.User
//
//	err := r.DB.Where("email = ? OR username = ?", email, username).First(&m).Error
//	if err != nil {
//		return nil, err
//	}
//
//	return &m, nil
//}

func (r *Users) GetByUsername(username string) (*models.User, error) {
	var m models.User

	err := r.DB.Where("username = ?", username).First(&m).Error
	if err != nil {
		return nil, err
	}

	return &m, nil
}

func (r *Users) ExistsByUsername(username string) bool {
	var exists bool

	r.DB.Model(&models.User{}).
		Select("count(*) > 0").
		Where("username = ?", username).
		Find(&exists)

	return exists
}

func (r *Users) ExistsByEmail(email string) bool {
	var exists bool

	r.DB.Model(&models.User{}).
		Select("count(*) > 0").
		Where("email = ?", email).
		Find(&exists)

	return exists
}
