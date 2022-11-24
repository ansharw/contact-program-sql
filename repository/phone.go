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

// func (repo *phoneRepository) Insert(ctx context.Context, phone model.Phone) (model.Phone, error) {
// 	var query string = "INSERT INTO phone(contact_id, phone) VALUES(?, ?)"
// 	_, phones := phone.GetPhone()
// 	res, err := repo.db.ExecContext(ctx, query, phones)
// 	if err != nil {
// 		return phone, err
// 	}
// 	lastInsertId, _ := res.LastInsertId()
// 	phone.SetPhone(int(lastInsertId), *phones)
// 	return phone, nil
// }

func (repo *phoneRepository) InsertPhones(ctx context.Context, phoneDatas []model.Phone, contact_id int) ([]model.Phone, error) {
	var query string = "INSERT INTO phone_data(contact_id, phone) VALUES(?, ?)"
	stmt, err := repo.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for _, v := range phoneDatas {
		_, phone := v.GetPhone()
		res, err := stmt.ExecContext(ctx, contact_id, phone)
		if err != nil {
			return nil, err
		}
		lastInsertId, _ := res.LastInsertId()
		v.SetPhone(int(lastInsertId), *phone)
	}
	return phoneDatas, nil
}

func (repo *phoneRepository) DeletePhoneByContactId(ctx context.Context, contact_id int) error {
	var query string = "DELETE FROM phone_data WHERE contact_id=?;"
	_, err := repo.db.ExecContext(ctx, query, contact_id)
	if err != nil {
		return err
	}
	return nil
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
