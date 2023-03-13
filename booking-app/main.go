package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"
const conferenceTickets = 50
var remainingTickets uint = conferenceTickets
var bookings = make([]User, 0)

type User struct {
	firstName string
	lastName string
	email string
	tickets uint
}

var waitGroup = sync.WaitGroup{}

func main() {

	greetUsers()

	// Infinite Loop
	// for {
		userData := getUserInputs()

		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(userData, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userData)
			waitGroup.Add(1)
			go sendTicket(userData)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				// break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address your entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	// }

	waitGroup.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInputs() User {
	var user User

	fmt.Println("Enter your first name: ")
	fmt.Scan(&user.firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&user.lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&user.email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&user.tickets)

	return user
}

func bookTicket(userData User) {
	remainingTickets -= userData.tickets

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v.\n", bookings)

	fmt.Printf(
		"Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n",
		userData.firstName, userData.lastName, userData.tickets, userData.email,
	)
	fmt.Printf("%v tickets remaining for \"%v\"\n", remainingTickets, conferenceName)
}

func sendTicket(user User)  {
	ticket := fmt.Sprintf("%v tickets for %v %v", user.tickets, user.firstName, user.lastName)
	fmt.Println("##########################################")
	fmt.Printf("Sending  %v tickets to %v...\n", ticket, user.firstName + " " + user.lastName)
	time.Sleep(30 * time.Second)
	fmt.Println("##########################################")
	fmt.Println("Done!")
	fmt.Println("")
	waitGroup.Done()
}
