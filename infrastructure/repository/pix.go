package repository

import (
	"fmt"

	"github.com/BahMartins/imersao-codepix/codepix-go/domain/model"
	"github.com/jinzhu/gorm"
)

type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

func (pixRepository PixKeyRepositoryDb) AddBank(bank *model.Bank) error {
	err := pixRepository.Db.Create(bank).Error

	if err != nil {
		return err
	}
	return nil
}

func (pixRepository PixKeyRepositoryDb) AddAccount(account *model.Account) error {
	err := pixRepository.Db.Create(account).Error

	if err != nil {
		return err
	}
	return nil
}

func (pixRepository PixKeyRepositoryDb) RegisterKey(pixKey *model.PixKey) (*model.PixKey, error) {
	err := pixRepository.Db.Create(pixKey).Error

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}

func (pixRepository PixKeyRepositoryDb) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	pixRepository.Db.Preload("Account.Bank").First(&pixKey, "kind = ? and key = ?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("no key was found")
	}

	return &pixKey, nil
}

func (pixRepository PixKeyRepositoryDb) FindAccount(id string) (*model.Account, error) {
	var account model.Account

	pixRepository.Db.Preload("Bank").First(&account, "id = ?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("no account found")
	}

	return &account, nil
}

func (pixRepository PixKeyRepositoryDb) FindBank(id string) (*model.Bank, error) {
	var bank model.Bank

	pixRepository.Db.First(&bank, "id = ?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("no bank found")
	}

	return &bank, nil
}
