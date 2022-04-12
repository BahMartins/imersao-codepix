package usecase

import (
	"errors"

	"github.com/BahMartins/imersao-codepix/codepix-go/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepository
}

func (pixUseCase *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := pixUseCase.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(kind, account, key)
	if err != nil {
		return nil, err
	}

	pixUseCase.PixKeyRepository.RegisterKey(pixKey)
	if pixKey.ID == "" {
		return nil, errors.New("unable to create new key at the moment")
	}

	return pixKey, nil
}

func (pixUseCase *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixkey, err := pixUseCase.PixKeyRepository.FindKeyByKind(key, kind)

	if err != nil {
		return nil, err
	}

	return pixkey, nil
}
