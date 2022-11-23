package template

import (
	"contact-program-fundamental/controller"
	"contact-program-fundamental/helper"
	"database/sql"
	"fmt"
)

type contactTemplate struct {
	contactHandler controller.ContactHandler
	db             *sql.DB
}

func NewContactTemplate(contactHandler controller.ContactHandler, db *sql.DB) *contactTemplate {
	return &contactTemplate{contactHandler, db}
}

func (c *contactTemplate) ListContact() {
	helper.ClearScreen()

	contacts, err := c.contactHandler.GetContacts()
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
	Menu(c.db)
}
