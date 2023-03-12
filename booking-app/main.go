package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets uint = conferenceTickets
	
	fmt.Printf("Conference tickets is %T, remainingTickets is %T \n", conferenceTickets, remainingTickets)

	fmt.Printf("Welcome to %v booking application.\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend!")

	var userName string = "Tom"
	var userTickets int = 2
	fmt.Printf("%v booked %v tickets.\n", userName, userTickets)
}
