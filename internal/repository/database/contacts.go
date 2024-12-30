package database

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"nootebook.com/internal/repository/boiler_models"
	"nootebook.com/internal/services/service_models/types"
)

type Contact interface {
	Get(ctx context.Context, name string) (*service_models.Contact, error)
	GetAll(ctx context.Context) ([]*service_models.Contact, error)
	Create(ctx context.Context, contact *service_models.Contact) error
	Delete(ctx context.Context, name string) error
	UpdateName(ctx context.Context, name string, updateNameParams string``) error
}

type ContactRepo struct {
	dbRead  *sql.DB
	dbWrite *sql.DB
	tx      *sql.Tx
}

func (cr *ContactRepo) Create(ctx context.Context, contact *service_models.Contact) error {
	var boilerContact boiler_models.Contact
	var boilerPhoneNumbers []*boiler_models.PhoneNumber
	boilerContact.Name = contact.Name

	for _, pn := range contact.PhoneNumbers {
		boilerPhoneNumbers = append(boilerPhoneNumbers, &boiler_models.PhoneNumber{
			Type:   pn.Type,
			Number: pn.Number,
		})
	}
	if err := boilerContact.Insert(ctx, exec(cr.dbWrite, cr.tx), boil.Infer()); err != nil {
		return err
	}

	return boilerContact.AddPhoneNumbers(ctx, exec(cr.dbWrite, cr.tx), true, boilerPhoneNumbers...)
}

func (cr *ContactRepo) GetAll(ctx context.Context) ([]*service_models.Contact, error) {
	var serviceContacts []*service_models.Contact
	var err error
	contacts, err := boiler_models.Contacts(qm.Load(boiler_models.ContactRels.PhoneNumbers)).All(ctx, exec(cr.dbRead, cr.tx))
	if err != nil {
		return nil, err
	}

	for _, contact := range contacts {
		var phoneNumbers []service_models.PhoneNumber
		for _, pn := range contact.R.PhoneNumbers {
			phoneNumbers = append(phoneNumbers, service_models.PhoneNumber{
				Type:   pn.Type,
				Number: pn.Number,
			})
		}
		serviceContacts = append(serviceContacts, &service_models.Contact{
			Name:         contact.Name,
			PhoneNumbers: phoneNumbers,
		})
	}
	return serviceContacts, nil
}

func (cr *ContactRepo) Delete(ctx context.Context, name string) error {
	_, err := boiler_models.Contacts(qm.Where("name = ?", name)).DeleteAll(ctx, exec(cr.dbWrite, cr.tx))
	if err != nil {
		return err
	}
	return nil
}

func (cr *ContactRepo) Get(ctx context.Context, name string) (*service_models.Contact, error) {
	serviceContact := new(service_models.Contact)
	var err error
	contact, err := boiler_models.Contacts(qm.Where("name = ?", name), qm.Load(boiler_models.ContactRels.PhoneNumbers)).One(ctx, exec(cr.dbRead, cr.tx))
	if err != nil {
		return nil, err
	}
	serviceContact.Name = contact.Name
	for _, pn := range contact.R.PhoneNumbers {
		serviceContact.PhoneNumbers = append(serviceContact.PhoneNumbers, service_models.PhoneNumber{
			Type:   pn.Type,
			Number: pn.Number,
		})
	}
	return serviceContact, nil
}

func (cr *ContactRepo) UpdateName(ctx context.Context, name string, updateNameParams string) error {
	c := &boiler_models.Contact{
		Name: name,
	}

	_, err := c.Update(ctx, exec(cr.dbWrite, cr.tx), boil.Whitelist(boiler_models.ContactColumns.Name))
	return err
}

func NewContactRepo(dbRead *sql.DB, dbWrite *sql.DB) *ContactRepo {
	return &ContactRepo{
		dbRead:  dbRead,
		dbWrite: dbWrite,
	}
}
