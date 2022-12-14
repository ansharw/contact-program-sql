package model

// "errors"
// "strconv"

type Contact struct {
	id         int
	name       string `required:"true"`
	phone_data []Phone
	email      string `type:"email"`
}

type Phone struct {
	id    int
	phone string
}

func (c *Contact) SetContact(id int, name string, phone_data []Phone, email string) {
	c.id = id
	c.name = name
	c.phone_data = phone_data
	c.email = email
}

func (c *Contact) UpdateContact(id int, name string, phone_data []Phone, email string) {
	c.id = id
	c.name = name
	c.phone_data = phone_data
	c.email = email
}

func (c *Contact) GetContact() (*int, *string, *[]Phone, *string) {
	return &c.id, &c.name, &c.phone_data, &c.email
}

func (c *Contact) GetId() *int {
	return &c.id
}

func (c *Contact) GetName() *string {
	return &c.name
}

func (c *Contact) GetEmail() *string {
	return &c.email
}

// buat banyak phone
func (p *Contact) SetPhone(phoneDatas []Phone) {
	p.phone_data = phoneDatas
}

func (p *Phone) SetPhone(id int, phone string) {
	p.id = id
	p.phone = phone
}

func (p *Phone) GetPhone() (*int, *string) {
	return &p.id, &p.phone
}

// var Contacts []Contact

// func GetLastId() int {
// 	if Contacts == nil {
// 		return 0
// 	} else {
// 		var tempId int
// 		for _, v := range Contacts {
// 			if tempId < v.id {
// 				tempId = v.id
// 			}
// 		}
// 		return tempId
// 	}
// }

// func SearchById(id *int) (Contact, error) {
// 	var contact Contact
// 	for _, v := range Contacts {
// 		if v.id == *id {
// 			return v, nil
// 		}
// 	}
// 	return contact, errors.New("data tidak ditemukan")
// }

// func GetIndex(id *int) (int, error) {
// 	for i, v := range Contacts {
// 		if v.id == *id {
// 			return i, nil
// 		}
// 	}
// 	return 0, errors.New("Id " + strconv.Itoa(*id) + " tidak ditemukan")
// }

// func (c *Contact) Add(datas ...map[string]interface{}) {
// 	for _, v := range datas {
// 		c.id = GetLastId() + 1
// 		c.name = v["name"].(string)
// 		c.phone = v["phone"].(string)
// 		c.email = v["email"].(string)
// 	}

// 	Contacts = append(Contacts, *c)
// }

// func (c *Contact) Edit(datas ...map[string]interface{}) {
// 	for _, v := range datas {
// 		c.name = v["name"].(string)
// 		c.phone = v["phone"].(string)
// 		c.email = v["email"].(string)
// 	}
// 	index, err := GetIndex(&c.id)
// 	if err == nil {
// 		Contacts[index].name = c.name
// 		Contacts[index].phone = c.phone
// 		Contacts[index].email = c.email
// 	}
// }

// func (c *Contact) Delete() {
// 	var index int
// 	for i, v := range Contacts {
// 		if v.id == c.id {
// 			index = i
// 		}
// 	}

// 	if index == len(Contacts)-1 {
// 		Contacts = Contacts[:index]
// 	} else if index == 0 {
// 		Contacts = Contacts[1:]
// 	} else {
// 		Contacts = append(Contacts[:index], Contacts[index+1:]...)
// 	}
// }

// func (c *Contact) GetFields() (int, string, string, string) {
// 	return c.id, c.name, c.phone, c.email
// }
