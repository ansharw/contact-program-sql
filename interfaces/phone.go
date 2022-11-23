package interfaces

import (
	"contact-program-fundamental/model"
	"context"
)

type PhoneInterface interface {
	GetPhoneByContactId(ctx context.Context, contact_id int) ([]model.Phone, error)
}