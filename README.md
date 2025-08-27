# CLI Contact Manager

A simple **command-line contact management application** written in Go.  
The app allows you to **add, list, search, edit, and delete contacts**, with **persistent storage using JSON files**.

---

## Features

- **Add contacts** (first name, last name, phone number, email)  
- **List all saved contacts**  
- **Search contacts** by first or last name (case-insensitive)  
- **Edit contacts** and save changes  
- **Delete contacts**  
- **Data persistence** using a JSON file (`contacts.json`)  

---

## Installation & Setup

### 1. Clone the repository
```bash
git clone https://github.com/meedaycodes/cli-contact-manager.git
cd cli-contact-manager

### 2. Build the project
go build -o contact-app main.go
