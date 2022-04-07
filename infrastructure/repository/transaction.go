package repository

import (
	"fmt"

	"github.com/BahMartins/imersao-codepix/codepix-go/domain/model"
	"github.com/jinzhu/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

func (transactionRepository *TransactionRepositoryDb) Registe(transaction *model.Transaction) error {
	err := transactionRepository.Db.Create(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (transactionRepository *TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := transactionRepository.Db.Save(transaction).Error

	if err != nil {
		return err
	}

	return nil
}

func (transactionRepository *TransactionRepositoryDb) Find(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	transactionRepository.Db.Preload("Account.Bank").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction was found")
	}

	return &transaction, nil
}
