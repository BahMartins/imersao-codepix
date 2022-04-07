package usecase

import (
	"errors"
	"log"

	"github.com/BahMartins/imersao-codepix/codepix-go/domain/model"
)

type TransactionUseCase struct {
	TransactionRepository model.TransactionRepository
	PixRepository         model.PixKeyRepository
}

func (transactionUseCase *TransactionUseCase) Register(accountId string, amount float64, pixKeyTo string, pixKeyKindTo string, description string) (*model.Transaction, error) {

	account, err := transactionUseCase.PixRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := transactionUseCase.PixRepository.FindKeyByKind(pixKeyTo, pixKeyKindTo)
	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account, amount, pixKey, description)
	if err != nil {
		return nil, err
	}

	transactionUseCase.TransactionRepository.Save(transaction)

	if transaction.ID != "" {
		return transaction, nil
	}

	return nil, errors.New("unable to process this transaction")
}

func (transactionUseCase *TransactionUseCase) Confirm(transactionId string) (*model.Transaction, error) {

	transaction, err := transactionUseCase.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}

	transaction.Status = model.TransactionConfirmed

	err = transactionUseCase.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (transactionUseCase *TransactionUseCase) Complete(transactionId string) (*model.Transaction, error) {

	trasaction, err := transactionUseCase.TransactionRepository.Find(transactionId)
	if err != nil {
		log.Println("Transaction not found", transactionId)
		return nil, err
	}

	trasaction.Status = model.TransactionCompleted
	err = transactionUseCase.TransactionRepository.Save(trasaction)
	if err != nil {
		return nil, err
	}

	return trasaction, nil
}

func (transactionUseCase *TransactionUseCase) Error(transactionId string, reason string) (*model.Transaction, error) {

	transaction, err := transactionUseCase.TransactionRepository.Find(transactionId)
	if err != nil {
		return nil, err
	}

	transaction.Status = model.TransactionError
	transaction.CancelDescription = reason

	err = transactionUseCase.TransactionRepository.Save(transaction)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
