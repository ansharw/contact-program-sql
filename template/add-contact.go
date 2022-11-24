package template

import (
	"contact-program-fundamental/helper"
	"contact-program-fundamental/model"
	"fmt"
	"reflect"
	"strings"
)

// // with return struct sehabis dari insert data
// func (c *contactTemplate) InsertContact() (model.Contact, error) {
// 	helper.ClearScreen()
// 	var email string
// 	var phoneSlice []string
// 	fmt.Println("Add Contact")
// 	fmt.Println("===============")
// 	name := InputName()
// 	phone := *InputPhone(&phoneSlice)
// 	fmt.Print("Email: ")
// 	fmt.Scanln(&email)

// 	var contact model.Contact
// 	res, err := c.contactHandler.InsertContact(name, email, phone)
// 	if err != nil {
// 		return contact, err
// 	}

// 	//Message berhasil
// 	fmt.Println("")
// 	fmt.Println("Data berhasil di input.")
// 	helper.BackHandler()
// 	Menu(c.db)
// 	return res, err
// }

// without return struct sehabis dari insert data
func (c *contactTemplate) InsertContact() error {
	helper.ClearScreen()
	var email string
	var phoneSlice []string
	fmt.Println("Add Contact")
	fmt.Println("===============")
	name := InputName()
	phone := *InputPhone(&phoneSlice)
	fmt.Print("Email: ")
	fmt.Scanln(&email)

	// var contact model.Contact
	err := c.contactHandler.InsertContact(name, email, phone)
	if err != nil {
		return err
	}

	//Message berhasil
	fmt.Println("")
	fmt.Println("Data berhasil di input.")
	helper.BackHandler()
	Menu(c.db)
	return err
}

func InputName() string {
	var name string
	fmt.Print("Name: ")
	fmt.Scanln(&name)

	if !ValidateName(&name) {
		fmt.Println("Name tidak boleh kosong")
		InputName()
	}
	return name
}

func InputPhone(phoneSlice *[]string) *[]string {
	var phone string
	fmt.Print("Phone: ")
	fmt.Scanln(&phone)
	*phoneSlice = append(*phoneSlice, phone)
	var lagi string
	fmt.Print("input phone lagi: ")
	fmt.Scanln(&lagi)
	if strings.ToLower(lagi) == "y" {
		InputPhone(phoneSlice)
	}
	return phoneSlice
}

func ValidateName(name *string) bool {
	var c model.Contact
	typeOf := reflect.TypeOf(c)
	if typeOf.Field(1).Tag.Get("required") == "true" {
		if *name == "" {
			return false
		}
	}
	return true
}
