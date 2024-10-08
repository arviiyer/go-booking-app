package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberofTickets uint
}

func main() {

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 {

		firstName, lastName, fullName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketCount := validateUserInput(firstName, lastName, email, userTickets)

		if isValidEmail && isValidName && isValidTicketCount {
			bookTicket(userTickets, fullName, email)

			firstNames := FirstNames()
			fmt.Printf("These are all our bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is sold out. Come back next year")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Invalid First Name or Last Name.")
			}
			if !isValidEmail {
				fmt.Println("Invalid Email")
			}
			if !isValidTicketCount {
				fmt.Println("Invalid Ticket Count")
			}
			println("Try again bozo")
			continue
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func FirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, string, uint) {
	var firstName string
	var lastName string
	var fullName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)

	fullName = firstName + " " + lastName

	fmt.Print("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, fullName, email, userTickets
}

func bookTicket(userTickets uint, fullName string, email string) {
	remainingTickets = remainingTickets - userTickets

	names := strings.Fields(fullName)
	firstName := names[0]
	lastName := names[1]

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberofTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", fullName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
