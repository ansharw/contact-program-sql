package controller

import (
	"contact-program-fundamental/interfaces"
	"contact-program-fundamental/model"
	"context"
)

type ContactHandler interface {
	GetContacts() ([]model.Contact, error)
	InsertContact(name, email string, phone []string) (model.Contact, error)
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

func (c *contacthandler) InsertContact(name, email string, phone []string) (model.Contact, error) {
	ctx := context.Background()

	var contact model.Contact
	var phoneDatas []model.Phone
	contact.SetContact(0, name, phoneDatas, email)
	
	res, err := c.contactInterface.Insert(ctx, contact)
	// c.phoneInterface.Insert(ctx, phoneDatas)

	return res, err
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
