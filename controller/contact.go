package controller

import (
	"contact-program-fundamental/interfaces"
	"contact-program-fundamental/model"
	"context"
)

type ContactHandler interface {
	GetContacts() ([]model.Contact, error)
	// with nilai balikan sehabis insert
	// InsertContact(name, email string, phone []string) (model.Contact, error)
	// without nilai balikan sehabis insert
	InsertContact(name, email string, phone []string) error
	DeleteContact(id int) error
	// SearchContactById(id int) (model.Contact, error)
	GetContact(id int) (model.Contact, error)
	UpdateContact(id int, name, email string, phone []string) error
}

type contacthandler struct {
	contactInterface interfaces.ContactInterface
	phoneInterface   interfaces.PhoneInterface
}

func NewContactHandler(contactInterface interfaces.ContactInterface, phoneInterface interfaces.PhoneInterface) *contacthandler {
	return &contacthandler{contactInterface, phoneInterface}
}

func (c *contacthandler) GetContacts() ([]model.Contact, error) {
	ctx := context.Background()

	contacts, err := c.contactInterface.FindAll(ctx)
	if err != nil {
		return contacts, err
	}

	for i, v := range contacts {
		id, name, _, email := v.GetContact()
		phoneDatas, err := c.phoneInterface.GetPhoneByContactId(ctx, *id)
		if err != nil {
			return nil, err
		}
		contacts[i].SetContact(*id, *name, phoneDatas, *email)
	}

	return contacts, nil
}

// // with return struct dari sehabis insert data
// func (c *contacthandler) InsertContact(name, email string, phone []string) (model.Contact, error) {
// 	ctx := context.Background()

// 	var contact model.Contact
// 	var phoneDatas []model.Phone

// 	// init struct contact from parameter
// 	contact.SetContact(0, name, phoneDatas, email)
// 	contact, err := c.contactInterface.Insert(ctx, contact)
// 	if err != nil {
// 		return contact, err
// 	}

// 	for _, v := range phone {
// 		var phone model.Phone
// 		phone.SetPhone(0, v)
// 		phoneDatas = append(phoneDatas, phone)
// 	}
// 	phoneDatas, err = c.phoneInterface.InsertPhones(ctx, phoneDatas, *contact.GetId())
// 	if err != nil {
// 		return contact, err
// 	}
// 	contact.SetPhone(phoneDatas)

// 	return contact, nil
// }

// without return struct dari sehabis insert data
func (c *contacthandler) InsertContact(name, email string, phone []string) error {
	ctx := context.Background()

	var contact model.Contact
	var phoneDatas []model.Phone

	// init struct contact from parameter
	contact.SetContact(0, name, phoneDatas, email)
	contact, err := c.contactInterface.Insert(ctx, contact)
	if err != nil {
		return err
	}

	for _, v := range phone {
		var phone model.Phone
		phone.SetPhone(0, v)
		phoneDatas = append(phoneDatas, phone)
	}
	phoneDatas, err = c.phoneInterface.InsertPhones(ctx, phoneDatas, *contact.GetId())
	if err != nil {
		return err
	}
	contact.SetPhone(phoneDatas)
	return nil
}

func (c *contacthandler) UpdateContact(id int, name, email string, phone []string) error {
	ctx := context.Background()
	var contact model.Contact
	var phoneDatas []model.Phone

	_, err := c.contactInterface.SearchById(ctx, id)
	if err == nil {
		err1 := c.phoneInterface.DeletePhoneByContactId(ctx, id)
		if err != nil {
			return err1
		}
		contact.UpdateContact(*contact.GetId(), name, phoneDatas, email)
		_, err := c.contactInterface.Update(ctx, contact)
		if err != nil {
			return err
		}
		for _, v := range phone {
			var phone model.Phone
			phone.SetPhone(*contact.GetId(), v)
			phoneDatas = append(phoneDatas, phone)
		}
		phoneDatas, err := c.phoneInterface.InsertPhones(ctx, phoneDatas, *contact.GetId())
		if err != nil {
			return err
		}
		contact.SetPhone(phoneDatas)
		return nil
	} else {
		return err
	}
}

func (c *contacthandler) DeleteContact(id int) error {
	ctx := context.Background()
	_, err := c.contactInterface.SearchById(ctx, id)
	if err == nil {
		err1 := c.phoneInterface.DeletePhoneByContactId(ctx, id)
		if err1 != nil {
			return err1
		}
		err := c.contactInterface.Delete(ctx, id)
		if err != nil {
			return err
		}
		return nil
	} else {
		return err
	}
}

// func (c *contacthandler) SearchContactById(id int) (model.Contact, error) {
// 	ctx := context.Background()
// 	contact, err := c.contactInterface.SearchById(ctx, id)
// 	if err != nil {
// 		return contact, err
// 	}
// 	return contact, nil
// }

func (c *contacthandler) GetContact(id int) (model.Contact, error) {
	ctx := context.Background()
	contact, err := c.contactInterface.SearchById(ctx, id)
	if err != nil {
		return contact, err
	}
	phoneDatas, err := c.phoneInterface.GetPhoneByContactId(ctx, id)
	if err != nil {
		return contact, err
	}
	contact.SetPhone(phoneDatas)
	return contact, nil
}

// func UpdateContactHandler(co interfaces.ContactInterface, name, phone, email string) {
// 	datas := map[string]interface{}{
// 		"name":  name,
// 		"phone": phone,
// 		"email": email,
// 	}
// 	co.Edit(datas)
// }

// func InsertContactHandler(co interfaces.ContactInterface, name, phone, email string) {
// 	datas := map[string]interface{}{
// 		"name":  name,
// 		"phone": phone,
// 		"email": email,
// 	}
// 	co.Add(datas)
// }

// func DeleteContactHandler(co interfaces.ContactInterface) {
// 	co.Delete()
// }

// func InsertCustomerHandler(co interfaces.ContactInterface, name string, age int) {
// 	datas := map[string]interface{}{
// 		"name": name,
// 		"age":  age,
// 	}
// 	co.Add(datas)
// }
