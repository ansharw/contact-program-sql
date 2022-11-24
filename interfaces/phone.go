package interfaces

import (
	"contact-program-fundamental/model"
	"context"
)

type PhoneInterface interface {
	GetPhoneByContactId(ctx context.Context, contact_id int) ([]model.Phone, error)
	// Insert(ctx context.Context, phone []model.Phone) (model.Phone, error)	
	InsertPhones(ctx context.Context, phoneDatas []model.Phone, contact_id int) ([]model.Phone, error)
	DeletePhoneByContactId(ctx context.Context, contact_id int) error
}