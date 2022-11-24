package template

import (
	"contact-program-fundamental/helper"
	"fmt"
	"strings"
)

func (c *contactTemplate) EditContact() {
	helper.ClearScreen()
	fmt.Println("Edit Contact")
	fmt.Println("===============")
	var search int
	fmt.Print("Input id yang akan di ubah: ")
	fmt.Scanln(&search)

	contact, err := c.contactHandler.GetContact(search)
	if err != nil {
		fmt.Println(err.Error())
		var jeda string
		fmt.Scanln(&jeda)
		c.EditContact()
	} else {
		_, name, phone, email := contact.GetContact()
		fmt.Println("Data Ditemukan")
		fmt.Println("===============")
		fmt.Println("Nama:", *name)
		fmt.Println("Phone:", helper.PhoneToString(*phone))
		fmt.Println("Email:", *email)
		fmt.Println("===============")
	}
	fmt.Println("")
	fmt.Println("")

	var name, email string
	var phoneSlice []string
	fmt.Println("Form Contact")
	fmt.Println("===============")
	fmt.Print("Name: ")
	fmt.Scanln(&name)
	UpdatePhone(&phoneSlice)
	fmt.Print("Email: ")
	fmt.Scanln(&email)

	err1 := c.contactHandler.UpdateContact(*contact.GetId(), name, email, phoneSlice)
	if err1 != nil {
		panic(err1)
	}

	//Message berhasil
	fmt.Println("")
	fmt.Println("Data berhasil di update.")
	helper.BackHandler()
	Menu(c.db)

	// contact, err := model.SearchById(&search)
	// if err != nil {
	// 	EditContact()
	// } else {
	// 	_, name, phone, email := contact.GetFields()
	// 	fmt.Println("Data Ditemukan")
	// 	fmt.Println("===============")
	// 	fmt.Println("Nama:", name)
	// 	fmt.Println("Phone:", phone)
	// 	fmt.Println("Email:", email)
	// 	fmt.Println("===============")
	// }

	// fmt.Println("")
	// fmt.Println("")

	// var name, phone, email string
	// fmt.Println("Form Contact")
	// fmt.Println("===============")
	// fmt.Print("Name: ")
	// fmt.Scanln(&name)
	// fmt.Print("Phone: ")
	// fmt.Scanln(&phone)
	// fmt.Print("Email: ")
	// fmt.Scanln(&email)
	// controller.UpdateContactHandler(&contact, name, phone, email)

	// //Message berhasil
	// fmt.Println("")
	// fmt.Println("Data berhasil di update.")
	// helper.BackHandler()
	// Menu()
}

func UpdatePhone(phoneSlice *[]string) {
	var phone string
	fmt.Print("Phone: ")
	fmt.Scanln(&phone)
	*phoneSlice = append(*phoneSlice, phone)
	var lagi string
	fmt.Print("input phone lagi (y/t): ")
	fmt.Scanln(&lagi)
	if strings.ToLower(lagi) == "y" {
		InputPhone(phoneSlice)
	}
}
