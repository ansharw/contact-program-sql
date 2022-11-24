package template

import (
	"contact-program-fundamental/helper"
	"fmt"
)

func (c *contactTemplate) DeleteContact() {
	helper.ClearScreen()
	fmt.Println("Delete Contact")
	fmt.Println("===============")
	var id int
	fmt.Print("Input id yang akan di delete: ")
	fmt.Scanln(&id)

	err := c.contactHandler.DeleteContact(id)
	if err != nil {
		panic(err)
	}

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

	// confirmDelete()
	// controller.DeleteContactHandler(&contact)

	//Message berhasil
	fmt.Println("")
	fmt.Println("Data berhasil di didelete.")
	helper.BackHandler()
	Menu(c.db)
}

// func confirmDelete() {
// 	fmt.Print("Apakah yakin ingin dihapus(y/t)")
// 	var confirm string
// 	fmt.Scanln(&confirm)
// 	switch confirm {
// 	case "y":
// 	case "t":
// 		Menu()
// 	default:
// 		confirmDelete()
// 	}
// }
