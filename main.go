package main

import (
	"fmt"
	"strings"
)

func main() {
	var confernceName = "Go Confernce"
	const conferenceTickets = 50
	var remainingTickets uint = 50
	bookings := []string{}

	fmt.Printf("welcome to our %v booking application", confernceName)
	fmt.Println("We have total ", conferenceTickets, "tickets and ", remainingTickets, "are still available")
	fmt.Println("Get your ticket here to attend")
	for {
		var firstName string
		var lastName string
		var email string
		var userTicket uint

		fmt.Println("enter your first name ")
		fmt.Scan(&firstName)

		fmt.Println("enter your last name ")
		fmt.Scan(&lastName)

		fmt.Println("enter your email  ")
		fmt.Scan(&email)

		fmt.Println("enter number of tickets")

		fmt.Scan(&userTicket)

		remainingTickets = remainingTickets - userTicket

		bookings = append(bookings, firstName+" "+lastName)
		fmt.Printf("Thank you %v %v for booking %v ticket , you will receive a confirmation email to you email address %v ", firstName, lastName, userTicket, email)
		fmt.Printf("there are %v tickets for %v", remainingTickets, confernceName)
		fmt.Printf("there are all booking %v", bookings)

		var firstNames []string

		for _, booking := range bookings {

			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("there are all booking %v", firstNames)

	}
}
