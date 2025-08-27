package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Contact struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

var contactList []Contact

func main() {

	loadContactsFromFile()

	fmt.Println("Welcome to the CLI Contact App")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1 - Add Contact")
		fmt.Println("2 - List Contacts")
		fmt.Println("3 - Search Contact")
		fmt.Println("4 - Edit Contact")
		fmt.Println("5 - Delete Contact")
		fmt.Println("6 - Exit")

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
			editContact()
		case "5":
			deleteContact()
		case "6":
			fmt.Println("Exiting the app. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please select 1, 2, 3, or 4.")
		}
	}
}

func addContacts(reader *bufio.Reader) Contact {

	var contactDetails Contact

	fmt.Println("\nEnter your contact details in the format:")
	fmt.Println("firstname, lastname, phonenumber, email")
	fmt.Print("Input: ")

	contactDetailsInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
	}
	contactDetailsInput = strings.TrimSpace(contactDetailsInput)

	contactParts := strings.Split(contactDetailsInput, ",")

	if len(contactParts) != 4 {
		fmt.Println("Not enough values provided")
	}

	// Trim spaces for each part
	contactDetails.FirstName = strings.TrimSpace(contactParts[0])
	contactDetails.LastName = strings.TrimSpace(contactParts[1])
	contactDetails.PhoneNumber = strings.TrimSpace(contactParts[2])
	contactDetails.Email = strings.TrimSpace(contactParts[3])

	contactList = append(contactList, contactDetails)
	saveContactsToFile()

	fmt.Printf("Contact added: %s %s | Phone: %s | Email: %s\n",
		contactDetails.FirstName, contactDetails.LastName, contactDetails.PhoneNumber, contactDetails.Email)

	return contactDetails
}

func listContacts() {

	if len(contactList) == 0 {
		fmt.Println("No contacts available. Please add some first.")
		return
	}

	fmt.Println("\nContact List:")
	for i, contact := range contactList {
		fmt.Printf("%d. %s %s | Phone: %s | Email: %s\n", i+1, contact.FirstName, contact.LastName, contact.PhoneNumber, contact.Email)
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
		if strings.ToLower(contact.FirstName) == nameInput || strings.ToLower(contact.LastName) == nameInput {
			fmt.Printf("Found: %s %s | Phone: %s | Email: %s\n",
				contact.FirstName, contact.LastName, contact.PhoneNumber, contact.Email)
			found = true
		}
	}

	if !found {
		fmt.Println("No contact found with that name.")
	}

}

func saveContactsToFile() {
	file, err := json.MarshalIndent(contactList, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling contacts:", err)
		return
	}
	err = os.WriteFile("contacts.json", file, 0644)
	if err != nil {
		fmt.Println("Error writing contacts file:", err)
	}
}

func loadContactsFromFile() {
	_, err := os.Stat("contacts.json")
	if os.IsNotExist(err) {
		return // File doesn't exist yet, skip loading
	}

	file, err := os.ReadFile("contacts.json")
	if err != nil {
		fmt.Println("Error reading contacts file:", err)
		return
	}

	err = json.Unmarshal(file, &contactList)
	if err != nil {
		fmt.Println("Error unmarshalling contacts:", err)
	}
}

func deleteContact() {
	var name string
	fmt.Println("Enter the firstname or lastname of the contact to delete:")
	fmt.Scan(&name)

	for i, c := range contactList {
		if c.FirstName == name || c.LastName == name {
			contactList = append(contactList[:i], contactList[i+1:]...)
			fmt.Println("Contact deleted.")
			saveContactsToFile()
			return
		}
	}
	fmt.Println("Contact not found.")
}

func editContact() {
	var name string
	fmt.Println("Enter the firstname or lastname of the contact to edit:")
	fmt.Scan(&name)

	for i, c := range contactList {
		if c.FirstName == name || c.LastName == name {
			fmt.Println("Enter new details: firstname, lastname, phone, email")
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			parts := strings.Split(input, ",")
			if len(parts) == 4 {
				contactList[i].FirstName = strings.TrimSpace(parts[0])
				contactList[i].LastName = strings.TrimSpace(parts[1])
				contactList[i].PhoneNumber = strings.TrimSpace(parts[2])
				contactList[i].Email = strings.TrimSpace(parts[3])
				fmt.Println("Contact updated.")
				saveContactsToFile()
			} else {
				fmt.Println("Invalid input format.")
			}
			return
		}
	}
	fmt.Println("Contact not found.")
}
