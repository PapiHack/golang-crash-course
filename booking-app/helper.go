package main

import "strings"

func validateUserInput(user User, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(user.firstName) >= 2 && len(user.lastName) >= 2
	isValidEmail := strings.Contains(user.email, "@")
	isValidTicketNumber := user.tickets > 0 && user.tickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
