package services

import (
	"context"

	service_models "nootebook.com/internal/services/service_models/types"
)

type Contact interface {
	Get(ctx context.Context, name string) (*service_models.Contact, error)
	GetAll(ctx context.Context) ([]*service_models.Contact, error)
	Insert(ctx context.Context, phonebook *service_models.Contact) error
	Update(ctx context.Context, update any) (*service_models.Contact, error)
	Delete(ctx context.Context, name *string) error
}

type ContactService struct {
	contactRepo any
}

func NewContactService(cRepo any) Contact {
	return &ContactService{
		contactRepo: cRepo,
	}
}

func (cs *ContactService) Get(ctx context.Context, name string) (*service_models.Contact, error) {
	// result, err := cs.Database()
	// if err != nil {
	// 	return nil, err
	// }
	// return result, nil
	return nil, nil
}

func (cs *ContactService) GetAll(ctx context.Context) ([]*service_models.Contact, error) {
	return nil, nil
}

func (cs *ContactService) Insert(ctx context.Context, phonebook *service_models.Contact) error {
	return nil
}

func (cs *ContactService) Update(ctx context.Context, update any) (*service_models.Contact, error) {
	return nil, nil
}

func (cs *ContactService) Delete(ctx context.Context, name *string) error {
	return nil
}
