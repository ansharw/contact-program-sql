package template

import (
	"contact-program-fundamental/helper"
	"contact-program-fundamental/repository"
	"fmt"
	"os"
)

func Menu(contactRepository *repository.ContactRepository) {
	helper.ClearScreen()
	fmt.Println("Menu")
	fmt.Println("=================")
	fmt.Println("1. List Contact")
	fmt.Println("2. Insert Contact")
	fmt.Println("3. Update Contact")
	fmt.Println("4. Delete Contact")
	fmt.Println("5. Exit")

	var menu int
	fmt.Print("Pilih menu: ")
	fmt.Scanln(&menu)

	switch menu {
	case 1:
		ListContact(contactRepository)
	// case 2:
	// 	AddContact()
	// case 3:
	// 	EditContact()
	// case 4:
	// 	DeleteContact()
	case 5:
		os.Exit(0)
	default:
		Menu(contactRepository)
	}
}
