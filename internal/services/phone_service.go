package services

import (
	"context"

	"nootebook.com/types"
)

type PhoneBookService interface {
	Get(ctx context.Context, id string) (*types.PhoneBook, error)
	GetAll(ctx context.Context) ([]*types.PhoneBook, error)
	Insert(ctx context.Context, phonebook *types.PhoneBook) error
	Update(ctx context.Context, filter any, update any) (*types.PhoneBook, error)
}
