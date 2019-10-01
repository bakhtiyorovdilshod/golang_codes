package main

import "testing"

var (
	cm *ContactManager

)

func TestContactManagerAdd (t *testing.T) {
	cm = NewContactManager()
	cm.add_list(Contact{first_name:"dilshod", last_name:"du", phone_number:"8766",age:21})
	if len(cm.contacts) == 0 {
		t.Error("Contact not added")
	}
}

func TestContactManagerUpdate (t *testing.T) {
	cm = NewContactManager()
	cm.add_list(Contact{first_name:"dilshod", last_name:"du", phone_number:"8766",age:21})
	x:=cm.list_contact()[0]
	cm.update_contact(0,Contact{first_name:"dilshodbek", last_name:"du", phone_number:"8766",age:21})
	if cm.list_contact()[0]==x {
		t.Error("Error")

	}
}

func TestContactManagerDelete (t *testing.T) {
	cm = NewContactManager()
	cm.add_list(Contact{first_name:"dilshod", last_name:"du", phone_number:"8766",age:21})
	my_contact_len:= len(cm.contacts)
	cm.delete_contact(0)
	check_contact_len:=len(cm.contacts)
	if my_contact_len == check_contact_len {
		t.Error("Error")

	}

}
