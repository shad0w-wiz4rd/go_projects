//basic cmd application: user registration and ticketing service (urts)

package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName string = "Cybersecurity Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketAmount := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketAmount {
			bookTicket(userTickets, firstName, lastName, email)

			wg.Add(1)
			go sendTicket(userTickets, firstName, lastName, email) // go
			//^Email Server takes 10-15 seconds to send user's  ticket
			//having another thread going will allow another person to sign up for the conference concurrently

			firstNames := printGuestList()
			fmt.Println("####################")
			fmt.Printf("First Names of Guests Booked: %v\n", firstNames)

			// end program
			if remainingTickets == 0 {
				fmt.Println("####################")
				fmt.Printf("Our %v is all booked. Come back next year.\n", conferenceName)
				fmt.Println("####################")
				break
			}
		} else {
			if !isValidName {
				fmt.Print("\nInvalid input: First and/or Last names. Try again\n")
			}
			if !isValidEmail {
				fmt.Print("Invalid input: Email. Try again\n")
			}
			if !isValidTicketAmount {
				fmt.Printf("Invalid input: You attempted to book %v tickets. There are only %v tickets remaining.\n", userTickets, remainingTickets)
			}
		}
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Println("####################")
	fmt.Printf("Welcome to our %v\n", conferenceName)
	fmt.Print("Book your tickets today!\n")
	fmt.Printf("Tickets Available: (%v/%v)\n", remainingTickets, conferenceTickets)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("\n####################")
	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter your email: ")
	fmt.Scan(&email)
	fmt.Print("Enter number of Tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, email, userTickets
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	var isValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidTicketAmount = (userTickets) > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketAmount
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//CREATE struct
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Println("####################")
	fmt.Printf("List of bookings: %v\n", bookings)
	fmt.Println("####################")
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receve a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Println("####################")
	fmt.Printf("%v: tickets remaining: (%v/%v)\n", conferenceName, remainingTickets, conferenceTickets)
}

func printGuestList() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(15 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("Sending ticket: %v\n \nCheck your email confirmation: %v\n", ticket, email)
	wg.Done()
}
