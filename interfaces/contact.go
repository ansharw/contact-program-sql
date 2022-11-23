package interfaces

import (
	"contact-program-fundamental/model"
	"context"
)

type ContactRepository interface {
	// Insert(ctx context.Context, comment model.Comment) (entity.Comment, error)
	// FindById(ctx context.Context, id int) (entity.Comment, error)
	FindAll(ctx context.Context) ([]model.Contact, error)
}
