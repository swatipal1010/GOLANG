package main //Give a package name to which this file beongs to

import (
	"fmt"
	"time"
)

// Introducing package-level variables(remove these variables from func. call and func. parameter list as they are accessible from here globally)
var conferenceName string = "Go Conference"

const conferenceTickets int = 50

var remainingTickets uint = 50

// var bookings [50]string								//Equivalent var bookings = [50]string{}
var bookings = make([]UserData, 0) //This is a empty slice of UserData struct. It doesn't have a fixed size rather it grows as per neeed.
														//Equ. to writing 'bookings := []map[string]string{}'

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

func main() {
	greetUsers()

	/*
		fmt.Printf("conferenceTicket is %T, remainingTickets is %T\n", conferenceTickets, remainingTickets)
		fmt.Println("Welcome to our", conferenceName, "booking application") //fmt is the package that contains Println method(Uppercase P in Println method indicates that it's a public method)
		fmt.Printf("We have %v tickets in total and %v tickets are still available!\n", conferenceTickets, remainingTickets)
		fmt.Println("Get your ticket to attend.")
	*/

	for { 													//Can mention condition after for loop 'for remainingTickets>0 && len(bookings)<=50'
		//Get user input
		firstName, lastName, email, userTicket := getUserInput()

		//Constraints (Multiple values returned by the function validUserInput) Function present in helper.go
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTicket)

		if isValidName && isValidEmail && isValidTicketNumber {
			//Book user ticket
			bookTicket(firstName, lastName, email, userTicket)
			//Send email ticket to user
			go sendTicket(firstName, lastName, email, userTicket)		//'go' adds concurrency by starting a goroutine for this func to execute

			//Use the function 'printFirtNames' to print firts names of all bookings
			firstNames := getFirstNames()
			fmt.Printf("The first name of all bookings are: %v\n", firstNames)

			var noTicketRemaining bool = remainingTickets == 0
			if noTicketRemaining {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Firstname or lastname you entered is too short. Try Again!!")
			}
			if !isValidEmail {
				fmt.Println("You entered an invalid email. It doesn't contains @ symbol.")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid.")
			}

		}

	}
}

// function to greet customers
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets in total and %v tickets are still available!\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket to attend.")
}

// Function to return firstname(slice of string values) of all the bookings
func getFirstNames() []string {
	firstNames := []string{} //slice to store only firstname of the customers

	for _, booking := range bookings { //We get 'index'(_) and 'value'(booking) while iterating through bookings slice using 'range' iterator

		//var names = strings.Fields(booking) //'names' is an array that contains 2 values(firstname & lastname) divided on whitespace seperator
		//using Fields method

		firstNames = append(firstNames, booking.firstName) //taking firstname from UserData struct
		fmt.Printf("List of bookings %v\n", bookings)
	}
	//fmt.Printf("The first name of all bookings are: %v\n",firstNames)
	return firstNames
}

// Function to get user input
func getUserInput() (string, string, string, uint) {
	var firstName string //When we don't assign a value to a variable immediately, we need to mention its datatype
	var lastName string
	var userTicket uint //uint allows for assigning postive integers only
	var email string

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want to buy: ")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket
}

// Function to book ticket
func bookTicket(firstName string, lastName string, email string, userTicket uint) {
	remainingTickets = remainingTickets - userTicket

	/*  bookings[0] = firstName+" "+lastName
	fmt.Printf("Whole array is:%v\n",bookings)
	fmt.Printf("First element of the array is : %v\n",bookings[0])
	fmt.Printf("Type of array is %T:\n",bookings)
	fmt.Printf("Length of the array is: %v\n",len(bookings))
	*/

	//Create map to store user data
	var userData = UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTicket,
	}

	//Adding elements in slice rather than in array
	bookings = append(bookings, userData)

	/*	fmt.Printf("Whole slice is:%v\n",bookings)
		fmt.Printf("First element of the slice is : %v\n",bookings[0])
		fmt.Printf("Type of slice is: %T\n",bookings)
		fmt.Printf("Length of the slice is: %v\n",len(bookings))
	*/

	fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTicket, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}

// Function to send ticket to the users who booked the conference tickets
func sendTicket(firstName string, lastName string, email string, userTicket uint) {
	time.Sleep(5 * time.Second)												//mimicing sleep of some time for background processing to take place
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicket, firstName, lastName) //Sprintf helps to save the string following it in a var
	fmt.Println("##############################################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("##############################################")
}
