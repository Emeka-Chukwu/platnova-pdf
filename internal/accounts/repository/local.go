package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	domain "pdf-go/domain/account"
)

const (
	open = "failed to open JSON file: %s"
	read = "failed to read JSON file: %s"
)

type AccountRepository interface {
	FetchAccountStatement(path string) ([]domain.AccountStatement, error)
}
type accountRepository struct {
}

func NewAccountRepository() AccountRepository {
	return &accountRepository{}
}

func (a *accountRepository) FetchAccountStatement(path string) ([]domain.AccountStatement, error) {
	result := []domain.AccountStatement{}
	file, err := os.Open(path)
	if err != nil {
		return result, errors.Join(err, fmt.Errorf(open, err))
	}
	defer file.Close()
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return result, errors.Join(err, fmt.Errorf(read, err))
	}
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return []domain.AccountStatement{}, err
	}
	return result, nil

}
