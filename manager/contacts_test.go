// Package manager models a simple contact list management system.
// Enables CRUD operations on contacts.
package manager

import (
	"slices"
	"testing"
)

func TestContact(t *testing.T){

	t.Run("Test contact", func(t *testing.T){
		contact := Contact{}

		got := contact.Name
		expect := ""

		if got != expect{
			t.Errorf("Got: %q, Expect: %q", got, expect)
		}
	})

	t.Run("Validate empty fields", func(t *testing.T){
		contact := Contact{Name: "", Email: "dakinola54@gmail.com", Phone: ""}
		list := ContactList{}

		err := list.AddContact(contact)

		if err == nil  {
			t.Errorf("expected an error but found none")
		}

	})

	t.Run("Validate invalid email address", func(t *testing.T){
		contact := Contact{Name: "David", Email: "dakinola54", Phone: "09077565788"}
		list := ContactList{}

		err := list.AddContact(contact)

		if err == nil  {
			t.Errorf("expected an error but found none")
		}

	})

	t.Run("Validate crrect phone number length", func(t *testing.T){
		contact := Contact{Name: "David", Email: "dakinola54@gmail.com", Phone: "090775657"}
		list := ContactList{}

		err := list.AddContact(contact)

		if err == nil  {
			t.Errorf("expected an error but found none")
		}

	})

	t.Run("Add a new contact", func(t *testing.T){
		contact := Contact{Name: "David", Email: "dakinola54@gmail.com", Phone: "09077565788"}
		list := ContactList{}

		err := list.AddContact(contact)

		if err != nil{
			t.Errorf("%q", err)
		}
		
	})
}

func TestContactList(t *testing.T){
	t.Run("Empty contact list", func(t *testing.T){
		list := ContactList{}

		got := list.GetAllContacts()
		expect := []Contact{}

		if !slices.Equal(got, expect){
			t.Errorf("Got: %v, Expect: %v", got, expect)
		}
	})

	t.Run("Get contact by email", func(t *testing.T){
		contact1 := Contact{Name: "David", Email: "dakinola54@gmail.com", Phone: "09077565788"}

		list := ContactList{}
		list.AddContact(contact1) 

		got, err := list.FindContactByEmail("dakinola54@gmail.com")
		expect := Contact{Name: "David", Email: "dakinola54@gmail.com", Phone: "09077565788"}

		if err == ErrContactNotFound{
			t.Errorf("%q", err)
			return 
		}

		if *got != expect{
			t.Errorf("Got: %q Expect: %q", got, expect)
		}

	})

	t.Run("All contacts", func(t *testing.T){
		contact1 := Contact{Name: "David", Email: "dakinola54@gmail.com", Phone: "09077565788"}
		contact2 := Contact{Name: "Ken", Email: "talktoken@gmail.com", Phone: "08169195589"}

		list := ContactList{}

		list.AddContact(contact1)
		list.AddContact(contact2)

		allContacts := list.GetAllContacts()
		expect := []Contact{contact1, contact2}

		if !slices.Equal(allContacts, expect){
			t.Errorf("Got: %v Expect: %v ", allContacts, expect)
		}

	})
}