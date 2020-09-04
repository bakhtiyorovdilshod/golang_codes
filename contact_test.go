package main

import "testing"

var (
	cm *ContactManager

)

func TestContactManagerAdd (t *testing.T) {
	cm = NewContacManager()
	cm.add_list(Contact{first_name:"dilshod", last_name:"du", phone_number:"8766",age:21})
	if len(cm.contacts)==0 {
		t.Error("Contact not added")
	}
	
}
func TestContactManagerUpdate (t *testing.T) {
	cm = NewContacManager()
	cm.add_list(Contact{first_name:"dilshod", last_name:"du", phone_number:"8766",age:21})
	x:=cm.list_contact()
	cm.update_list(0,Contact{first_name:"dilshodbek", last_name:"du", phone_number:"8766",age:21})
	if cm.list_contact()!=x {
		t.Error("Error")

	}
}