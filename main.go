package main

import (
	"contact-program-fundamental/database"
	"contact-program-fundamental/template"
)

func main() {
	db := database.GetConnection()
	// contactRepository := repository.NewContactRepository(db)
	template.Menu(db)
}
