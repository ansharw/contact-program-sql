package controller

import (
	"contact-program-fundamental/interfaces"
	"contact-program-fundamental/model"
	"context"
)

func GetContacts(contactInterface interfaces.ContactRepository) ([]model.Contact, error) {
	ctx := context.Background()
	return contactInterface.FindAll(ctx)
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
