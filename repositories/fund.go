package repositories

import (
	"holyways/models"

	"gorm.io/gorm"
)

type FundRepository interface {
	FindFund() ([]models.Fund, error)
	GetFund(ID int) (models.Fund, error)
	// GetFundStatus(ID int) (models.Transaction, error)
	CreateFund(fund models.Fund) (models.Fund, error)
	UpdateFund(fund models.Fund) (models.Fund, error)
	DeleteFund(fund models.Fund) (models.Fund, error)
}

func RepositoryFund(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindFund() ([]models.Fund, error) {
	var funds []models.Fund
	err := r.db.Preload("User").Preload("Transaction").Find(&funds).Error

	return funds, err
}

func (r *repository) GetFund(ID int) (models.Fund, error) {
	var fund models.Fund
	err := r.db.Preload("User").Preload("Transaction").Preload("Transaction.User").Preload("Transaction.Fund").First(&fund, ID).Error

	return fund, err
}

func (r *repository) CreateFund(fund models.Fund) (models.Fund, error) {
	err := r.db.Preload("User").Create(&fund).Error

	return fund, err
}

func (r *repository) UpdateFund(fund models.Fund) (models.Fund, error) {
	err := r.db.Save(&fund).Error

	return fund, err
}

func (r *repository) DeleteFund(fund models.Fund) (models.Fund, error) {
	err := r.db.Delete(&fund).Error

	return fund, err
}

// func (r *repository) GetFundStatus(ID int) (models.Transaction, error) {
// 	var trans models.Transaction
// 	err := r.db.Find(&trans, "id =?", ID).Error

// 	return trans, err
// }
