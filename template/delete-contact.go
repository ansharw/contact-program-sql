package template

import (
	"contact-program-fundamental/helper"
	"database/sql"
	"fmt"
)

func (c *contactTemplate) DeleteContact() {
	helper.ClearScreen()
	fmt.Println("Delete Contact")
	fmt.Println("===============")
	var search int
	fmt.Print("Input id yang akan di delete: ")
	fmt.Scanln(&search)

	contact, err := c.contactHandler.GetContact(search)
	if err != nil {
		fmt.Println(err.Error())
		var jeda string
		fmt.Scanln(&jeda)
		c.DeleteContact()
	} else {
		_, name, phone, email := contact.GetContact()
		fmt.Println("Data Ditemukan")
		fmt.Println("===============")
		fmt.Println("Nama:", *name)
		fmt.Println("Phone:", helper.PhoneToString(*phone))
		fmt.Println("Email:", *email)
		fmt.Println("===============")
	}
	confirmDelete(c.db)
	c.contactHandler.DeleteContact(search)
	//Message berhasil
	fmt.Println("Data berhasil di didelete.")
	helper.BackHandler()
	Menu(c.db)

	// err := c.contactHandler.DeleteContact(id)
	// if err != nil {
	// 	// Message gagal di delete
	// 	fmt.Println("")
	// 	fmt.Println("Data gagal di didelete.")
	// 	helper.BackHandler()
	// 	Menu(c.db)
	// }

	// contact, err := model.SearchById(&id)
	// if err != nil {
	// 	DeleteContact()
	// } else {
	// 	_, name, phone, email := contact.GetFields()
	// 	fmt.Println("Data Ditemukan")
	// 	fmt.Println("===============")
	// 	fmt.Println("Nama:", name)
	// 	fmt.Println("Phone:", phone)
	// 	fmt.Println("Email:", email)
	// 	fmt.Println("===============")
	// }
}

func confirmDelete(db *sql.DB) {
	fmt.Print("Apakah yakin ingin dihapus(y/t)")
	var confirm string
	fmt.Scanln(&confirm)
	switch confirm {
	case "y":
	case "t":
		Menu(db)
	default:
		confirmDelete(db)
	}
}
