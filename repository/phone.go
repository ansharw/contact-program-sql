package repository

import (
	"contact-program-fundamental/model"
	"context"
	"database/sql"
)

type phoneRepository struct {
	db *sql.DB
}

func NewPhoneRepository(db *sql.DB) *phoneRepository {
	return &phoneRepository{db}
}

func (repo *phoneRepository) GetPhoneByContactId(ctx context.Context, contact_id int) ([]model.Phone, error) {
	var phoneDatas []model.Phone
	var query string = "SELECT id, phone FROM phone_data WHERE contact_id = ?"
	rows, err := repo.db.QueryContext(ctx, query, contact_id)
	if err != nil {
		return phoneDatas, err
	}

	for rows.Next() {
		var phoneStruct model.Phone
		id, phone := phoneStruct.GetPhone()
		rows.Scan(id, phone)
		phoneDatas = append(phoneDatas, phoneStruct)
	}
	return phoneDatas, err
}
