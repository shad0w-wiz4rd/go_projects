package main

import (
	"fmt"
	"regexp"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Email     string
	Birthday  string
	Age       int
	Gender    string
	Ethnicity string
}

func main() {
	user, err := getUserInput()
	if err != nil {
		fmt.Println("Error getting user input:", err)
		return
	}
	fmt.Println("\nUser Information:")
	fmt.Println("First Name:", user.FirstName)
	fmt.Println("Last Name:", user.LastName)
	fmt.Println("Email:", user.Email)
	fmt.Println("Birthday:", user.Birthday)
	fmt.Println("Age:", user.Age)
	fmt.Println("Gender:", user.Gender)
	fmt.Println("Ethnicity:", user.Ethnicity)
}

func getUserInput() (User, error) {
	var firstName string
	var lastName string
	var email string
	var birthday string
	var gender string
	var ethnicity string

	// Prompt the user for each input field
	fmt.Print("Enter your first name: ")
	fmt.Scanln(&firstName)

	// Validate first name
	if !isValidStringInput(firstName) {
		fmt.Println("Invalid Input: Your first name should not contain special characters.")
		return getUserInput()
	}

	fmt.Print("Enter your last name: ")
	fmt.Scanln(&lastName)

	// Validate last name
	if !isValidStringInput(lastName) {
		fmt.Println("Invalid Input: Your last name should not contain special characters.")
		return getUserInput()
	}

	fmt.Print("Enter your email: ")
	fmt.Scanln(&email)

	// Validate email
	if !isValidEmail(email) {
		fmt.Println("Invalid Input. Your email address should only contain letters, numbers, and select special characters")
		return getUserInput()
	}

	fmt.Print("Enter your birthday (YYYY-MM-DD): ")
	fmt.Scanln(&birthday)

	// Calculate and validate age
	age, isValidAge := calculateAndValidateAge(birthday)
	if !isValidAge {
		fmt.Println("Invalid Input: Please enter a valid age (YYYY-MM-DD).")
		return getUserInput()
	}

	fmt.Print("Enter your gender (M, F, NB, Other): ")
	fmt.Scanln(&gender)

	if !isValidGenderInput(gender) {
		fmt.Println("Invalid Input: Please choose from one of the options listed.")
		return getUserInput()
	}

	fmt.Print("Enter your ethnicity: ")
	fmt.Scanln(&ethnicity)

	// Validate ethnicity
	if !isValidStringInput(ethnicity) {
		fmt.Println("Invalid Input: Please enter your ethnicity.")
		return getUserInput()
	}

	// Create and return a User object
	user := User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Birthday:  birthday,
		Age:       age,
		Gender:    gender,
		Ethnicity: ethnicity,
	}
	return user, nil
}

func isValidStringInput(name string) bool {
	// Regular expression to check if the name contains only letters and spaces
	regex := regexp.MustCompile("^[a-zA-Z\\s]+$")
	return regex.MatchString(name) && len(name) >= 2
}

func isValidEmail(email string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9@.+_-]+@[a-zA-Z0-9-._]+$`)
	return regex.MatchString(email)
}

func calculateAndValidateAge(birthday string) (int, bool) {
	today := time.Now()
	parsedBirthday, err := time.Parse("2006-01-02", birthday)
	if err != nil {
		fmt.Println("Invalid Input: Please enter a valid age (YYYY-MM-DD).") // Handle error appropriately
	}

	age := today.Year() - parsedBirthday.Year()

	if today.Month() < parsedBirthday.Month() || (today.Month() == parsedBirthday.Month() && today.Day() < parsedBirthday.Day()) {
		age-- // Adjust age if the person hasn't had their birthday yet this year
	}

	isValidAge := false

	if age < 18 {
		fmt.Println("Age Restriction: You must be at least 18 years old to proceed.")
	} else if age >= 125 {
		fmt.Println("Invalid Input: Please enter a valid age (YYYY-MM-DD).")
	} else {
		isValidAge = true
	}

	return age, isValidAge
}

func isValidGenderInput(gender string) bool {
	validGenders := []string{"M", "F", "NB", "Other"}
	for _, validGender := range validGenders {
		if validGender == gender {
			return true
		}
	}
	return false
}
