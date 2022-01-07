package helper
import (
	"strings"
)

func ValidateUserInput(userName string,middleName string,lastName string,email string,
	userTickets int,remainingTicket int) (bool,bool,bool){

	isValidName := len(userName) >= 2 && len(middleName) >= 2 && len(lastName) >=2
	isValidEmail := strings.Contains(email,"@")
	isValidTicketNumbers := userTickets > 0 && userTickets <= remainingTicket
	return isValidName ,isValidEmail,isValidTicketNumbers
}