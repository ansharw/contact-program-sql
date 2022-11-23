package template

import (
	"contact-program-fundamental/controller"
	"contact-program-fundamental/helper"
	"contact-program-fundamental/repository"
	"fmt"
)

func ListContact(contactRepository *repository.ContactRepository) {
	helper.ClearScreen()

	contacts, err := controller.GetContacts(contactRepository)
	if err != nil {
		panic(err)
	}
	fmt.Println("==========================================================")
	fmt.Println("ID\tNama\t\tPhone\t\tEmail")
	fmt.Println("==========================================================")
	if len(contacts) == 0 {
		fmt.Println("Data kosong")
	} else {
		for _, v := range contacts {
			id, name, _, email := v.GetContact()
			fmt.Printf("%v\t%v\t\t%v\n", *id, *name, *email)
		}
	}
	fmt.Println("==========================================================")
	helper.BackHandler()
	Menu(contactRepository)
}
