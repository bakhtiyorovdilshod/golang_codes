package main 

import (
	"fmt"
)
type Contact struct {
	first_name string
	last_name string
	age int
	phone_number string

}


type ContactManager struct {
	contacts []Contact
}

func NewContactManager() *ContactManager {
 	cm := ContactManager{}
 	cm.contacts = []Contact{}

 	return &cm
 }

func (contact *ContactManager) add_list(item Contact) {
    contact.contacts = append(contact.contacts, item)
}

func (contact *ContactManager) list_contact() []Contact{
	return contact.contacts
}

func (contact * ContactManager) update_contact(id int, c Contact) {
	contact.contacts[id] = c
}

func (contact * ContactManager) delete_contact(id int) {
	var i int 
	i = id
	contact.contacts = append(contact.contacts[:i], contact.contacts[i+1:]...)

}

func main() {
	cm := NewContactManager()
	cm.add_list(Contact{first_name:"dilshod", last_name:"du", phone_number:"8766",age:21})
	x:=cm.list_contact()[0]

	cm.update_contact(0,Contact{first_name:"dilshodbek", last_name:"du", phone_number:"8766",age:21})
    if x!=cm.list_contact()[0] {
    	fmt.Println("E")
    }
    fmt.Println(cm.list_contact()[0])
   

}


