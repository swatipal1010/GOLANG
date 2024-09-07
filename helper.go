// this file contains helper functions for the main application
package main

import "strings"

//variables and functions defined outside any function, can be accessed in all other files within the same package

//Function to validate user input. In the below line (bool,bool,bool) is the return type
func validateUserInput(firstName string, lastName string, email string, userTicket uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTicket > 0 && userTicket <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber 						//Mutiple return values are possible in GOLANG
}
