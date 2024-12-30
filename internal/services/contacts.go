package services

import (
	"context"
	"nootebook.com/internal/repository/database"
	service_models "nootebook.com/internal/services/service_models/types"
)

type Contact interface {
	Get(ctx context.Context, name string) (*service_models.Contact, error)
	GetAll(ctx context.Context) ([]*service_models.Contact, error)
	Insert(ctx context.Context, phonebook *service_models.Contact) error
	UpdateName(ctx context.Context, name string, updateNameParams string) error
	Delete(ctx context.Context, name string) error
}

type ContactService struct {
	contactRepo *database.ContactRepo
}

func NewContactService(cRepo *database.ContactRepo) *ContactService {
	return &ContactService{
		contactRepo: cRepo,
	}
}

func (cs *ContactService) Get(ctx context.Context, name string) (*service_models.Contact, error) {
	return cs.contactRepo.Get(ctx, name)
}

func (cs *ContactService) GetAll(ctx context.Context) ([]*service_models.Contact, error) {
	return cs.contactRepo.GetAll(ctx)
}

func (cs *ContactService) Insert(ctx context.Context, contact *service_models.Contact) error {
	return cs.contactRepo.Create(ctx, contact)
}

func (cs *ContactService) UpdateName(ctx context.Context, name string, updateNameParams string) error {
	return cs.contactRepo.UpdateName(ctx, name, updateNameParams)
}

func (cs *ContactService) Delete(ctx context.Context, name string) error {
	return cs.contactRepo.Delete(ctx, name)
}
