package repository

import (
	"contact-program-fundamental/model"
	"context"
	"database/sql"
)

type ContactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) *ContactRepository {
	return &ContactRepository{db}
}

func (repo *ContactRepository) FindAll(ctx context.Context) ([]model.Contact, error) {
	var query string = "SELECT id, name, email FROM contact"
	var contacts []model.Contact

	rows, err := repo.db.QueryContext(ctx, query)

	if err != nil {
		return contacts, err
	}

	for rows.Next() {
		var contact model.Contact
		id, name, _, email := contact.GetContact()
		rows.Scan(id, name, email)
		contacts = append(contacts, contact)
	}
	return contacts, err
}
