package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	firstName   string
	lastName    string
	phoneNumber string
	email       string
}

var contactList []Contact

func main() {
	fmt.Println("Welcome to the CLI Contact App")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1 - Add Contact")
		fmt.Println("2 - List Contacts")
		fmt.Println("3 - Search Contact")
		fmt.Println("4 - Exit")

		fmt.Print("Enter choice: ")
		choiceInput, _ := reader.ReadString('\n')
		choiceInput = strings.TrimSpace(choiceInput)

		switch choiceInput {
		case "1":
			addContacts(reader)
		case "2":
			listContacts()
		case "3":
			searchContact(reader)
		case "4":
			fmt.Println("Exiting the app. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select 1, 2, 3, or 4.")
		}
	}
}

func addContacts(reader *bufio.Reader) Contact {

	var contactDetails Contact

	for {

		fmt.Println("\nEnter your contact details in the format:")
		fmt.Println("firstname, lastname, phonenumber, email")
		fmt.Print("Input: ")

		contactDetailsInput, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		contactDetailsInput = strings.TrimSpace(contactDetailsInput)

		contactParts := strings.Split(contactDetailsInput, ",")

		if len(contactParts) != 4 {
			fmt.Println("Not enough values provided")
			continue
		}

		// Trim spaces for each part
		contactDetails.firstName = strings.TrimSpace(contactParts[0])
		contactDetails.lastName = strings.TrimSpace(contactParts[1])
		contactDetails.phoneNumber = strings.TrimSpace(contactParts[2])
		contactDetails.email = strings.TrimSpace(contactParts[3])

		break
	}

	contactList = append(contactList, contactDetails)

	fmt.Printf("Contact added: %s %s | Phone: %s | Email: %s\n",
		contactDetails.firstName, contactDetails.lastName, contactDetails.phoneNumber, contactDetails.email)

	return contactDetails
}

func listContacts() {

	if len(contactList) == 0 {
		fmt.Println("No contacts available. Please add some first.")
		return
	}

	fmt.Println("\nContact List:")
	for i, contact := range contactList {
		fmt.Printf("%d. %s %s | Phone: %s | Email: %s\n", i+1, contact.firstName, contact.lastName, contact.phoneNumber, contact.email)
	}

}

func searchContact(reader *bufio.Reader) {

	if len(contactList) == 0 {
		fmt.Println("Cannot search in an empty contact list.")
		return
	}

	fmt.Print("Enter first or last name to search: ")
	nameInput, _ := reader.ReadString('\n')
	nameInput = strings.TrimSpace(strings.ToLower(nameInput))

	found := false
	for _, contact := range contactList {
		if strings.ToLower(contact.firstName) == nameInput || strings.ToLower(contact.lastName) == nameInput {
			fmt.Printf("Found: %s %s | Phone: %s | Email: %s\n",
				contact.firstName, contact.lastName, contact.phoneNumber, contact.email)
			found = true
		}
	}

	if !found {
		fmt.Println("No contact found with that name.")
	}

}
