// Package manager models a simple contact list management system.
// Enables CRUD operations on contacts.
package manager

import (
	"errors"
	"strings"
)

// Contact defines the blueprint for a contact.
// A contact will have a name, an email address and a phone number.
type Contact struct{
	Name string
	Email string
	Phone string
}

// ContactList holds a list of contacts and the total number of contacts.
type ContactList struct{
	Contacts []Contact // slice of type "Contact".
	Total int
}

// AddContact method creates a new contact in the ContactList struct.
// It uses a pointer receiver because it modifies the Contacts field of the ContactList.
// It returns an error if any validation breaks during the creation of a contact.
func (cl *ContactList) AddContact(contact Contact) error{
	// Checks for empty fields in the contact information.
	if contact.Name == "" || contact.Email == "" || contact.Phone == ""{
		return errors.New("invalid fields")
	}
	// Checks whether the email field has a valid email address.
	if !strings.HasSuffix(contact.Email, "@gmail.com"){
		return errors.New("invalid email address")
	}
	// Checks whether the phone number has a value less than 10 or greater than 11.
	if len(contact.Phone) > 11 || len(contact.Phone) < 10{
		return errors.New("invalid phone number length (>=10 n <=11)")
	}

	cl.Contacts = append(cl.Contacts, contact)

	return nil
}

func (cl *ContactList) GetAllContacts() []Contact{
	return cl.Contacts
}