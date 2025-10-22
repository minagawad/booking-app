package main

import "fmt"

func main() {
	var confernceName = "Go Confernce"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("welcome to our %v booking application", confernceName)
	fmt.Println("We have total ", conferenceTickets, "tickets and ", remainingTickets, "are still available")

	fmt.Println("Get your ticket here to attend")

	var firstName string
	var lastName string
	var email string
	var userTicket int

	fmt.Println("enter your first name ")
	fmt.Scan(&firstName)

	fmt.Println("enter your last name ")
	fmt.Scan(&lastName)

	fmt.Println("enter your email  ")
	fmt.Scan(&email)

	fmt.Println("enter number of tickets")

	fmt.Scan(&userTicket)

	fmt.Printf("Thank you %v %v for booking %v ticket , you will receive a confirmation email to you email address %v ", firstName, lastName, userTicket, email)

}
