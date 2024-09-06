package main //Give a package name to which this file beongs to

import "fmt"

func main() {
	//Introducing variable
	var conferenceName string = "Go Conference"
	const tickets = 50
	remainingTickets := 40
	fmt.Println("Welcome to our",conferenceName,"booking application") //fmt is the package that contains Println method(Uppercase P in Println method indicates that it's a public method)
	fmt.Printf("We have %v tickets in total and %v tickets are still available!\n",tickets, remainingTickets)
	fmt.Println("Get your ticket to attend.")


	var username string			//When we don't assign a value to a variable immediately, we need to mention its datatype
	var userTicket uint			//uint allows for assigning postive integers only
	//ask for username
	username = "John"
	userTicket = 10
	fmt.Println(username)
	fmt.Println(userTicket)
																
	
}
