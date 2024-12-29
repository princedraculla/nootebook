package repository

import (
	"golang.org/x/net/context"
	service_models "nootebook.com/internal/services/service_models/types"
)

type Contact interface {
	Get(ctx context.Context, name string) (*service_models.Contact, error)
	GetAll(ctx context.Context) ([]*service_models.Contact, error)
	Create(ctx context.Context, contact *service_models.Contact) error
	Delete(ctx context.Context, name string) error
	Update(ctx context.Context, name string, contact *service_models.Contact) error
}

type contactRepo struct {
}
