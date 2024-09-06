package main //Give a package name to which this file beongs to

import "fmt"

func main() {
	//Introducing variable
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50	
	var bookings [50]string									//Equivalent var bookings = [50]string{}


	fmt.Printf("conferenceTicket is %T, remainingTickets is %T\n", conferenceTickets, remainingTickets)
	fmt.Println("Welcome to our", conferenceName, "booking application") //fmt is the package that contains Println method(Uppercase P in Println method indicates that it's a public method)
	fmt.Printf("We have %v tickets in total and %v tickets are still available!\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your ticket to attend.")

	
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

	remainingTickets = remainingTickets-userTicket

	bookings[0] = firstName+" "+lastName
	fmt.Printf("Whole array is:%v\n",bookings)
	fmt.Printf("First element of the array is : %v\n",bookings[0])
	fmt.Printf("Type of array is %T:\n",bookings)
	fmt.Printf("Length of the array is: %v\n",len(bookings))


	fmt.Printf("Thank You %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName,lastName, userTicket,email)
	fmt.Printf("%v tickets remaining for %v\n",remainingTickets,conferenceName)

}
