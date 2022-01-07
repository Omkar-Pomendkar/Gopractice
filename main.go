package main

import (
	"Gopractice/helper"
	"fmt"
	"strings"
	
)
	
var conferenceName = "Go Conference"
const conferenceTicket int = 50
var remainingTicket int = 50
var bookings = [] string{}
func main(){

	// Variables
	// var confernceName = "Go Conference"
	
	greetUsers()
	

	// fmt.Printf("conference name is %T conferenceTicket is %T remainingTicket is %T\n ",conferenceName,conferenceTicket,remainingTicket)
	// fmt.Printf("Welcome to %v Booking System\n",conferenceName);
	
	// fmt.Println("Get your Tickets here to attend");
	
	// fmt.Println(conferenceTicket)
	// fmt.Println(&remainingTicket)

	for{
		
		userName,middleName,lastName,email,userTickets := UserInput()
		isValidName,isValidEmail,isValidTicketNumbers := helper.ValidateUserInput(userName ,middleName ,lastName ,email ,userTickets,remainingTicket )

	
		// ask user to enter their name
		// userName = "Tom"

		// var bookings [50] string

		// bookings [0] = userName + " " + lastName

		
		
		

		// fmt.Printf("The Whole Array: %v \n",bookings)
		// fmt.Printf("The Whole Array: %v \n",bookings[0])
		// fmt.Printf("Type of Array: %T \n",bookings)
		// fmt.Printf("Length of Array: %v \n",len(bookings))

		if isValidEmail && isValidName && isValidTicketNumbers {

			bookTicket(userTickets ,userName,middleName ,lastName,email )
		
		// for each loop Iterating over List
		firstNames := getFirstName()
		fmt.Printf("The first Name of Bookings are %v ",firstNames)

		if remainingTicket == 0 {
			// end the Program
			fmt.Println("Confernce is Booked up come back next Year")
			break
			
		}
	}else {
			if !isValidName{
				println("Your name is to Short")
			}
			if !isValidEmail{
				println("Email id is Invalid")
			}
			if !isValidTicketNumbers{
				println("Ticket number you enter is wrong")
			}
			
		fmt.Println("Invalid Data Please try again")
		
	}
}
	
}

func greetUsers(){
	fmt.Printf("Welcome to %v ",conferenceName)
	fmt.Printf("We have this much total %v tickets %v are still remaining \n",conferenceTicket,remainingTicket)
}
func getFirstName() [] string  {

	firstNames := [] string {}  // SLice Variable
	for _, booking := range bookings {

		var names = strings.Fields(booking)
		firstNames = append(firstNames,names[0])
	}
	return firstNames
}


func UserInput() (string,string,string,string,int){
	var userName string
		var middleName string
		var lastName string
		var email string
		var userTickets int

		
		
		fmt.Println("Enter your first Name")
		fmt.Scan(&userName)

		fmt.Println("Enter your middle Name")
		fmt.Scan(&middleName)
		
		fmt.Println("Enter your last Name")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email Name")
		fmt.Scan(&email)

		fmt.Println("How many Tickets you want to Buy")
		fmt.Scan(&userTickets)
		
		return userName,middleName,lastName,email,userTickets
}
func bookTicket(userTickets int,userName string,middleName string,lastName string,email string){
		
	remainingTicket = remainingTicket-userTickets
	// fmt.Println(remainingTicket)
	bookings = append(bookings,userName + " " + lastName)
		


	fmt.Printf("Hello your Name is %v %v %v \n Your E-mail id is %v \n You have Booked %v Tickets", userName,middleName,lastName,email,userTickets)
	fmt.Printf(" \n %v Tickets are Remaining %v Confernece",remainingTicket,conferenceName)
}