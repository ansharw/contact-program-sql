package interfaces

import (
	"contact-program-fundamental/model"
	"context"
)

type ContactInterface interface {
	// FindById(ctx context.Context, id int) (entity.Comment, error)
	FindAll(ctx context.Context) ([]model.Contact, error)
	Insert(ctx context.Context, contact model.Contact) (model.Contact, error)
	Delete(ctx context.Context, id int) error
	SearchById(ctx context.Context, id int) (model.Contact, error)
}
