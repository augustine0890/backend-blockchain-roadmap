package main

import (
	"fmt"
	"pay/payment"
)
func main() {
	credit := payment.CreateCreditAccount(
		"Debra Ingram",
		"1111-2222-3333-4444",
		5,
		2025,
		123,
	)

	fmt.Printf("Owner name: %s\n", credit.OwnerName())
	fmt.Printf("Owner name: %s\n", credit.CardNumber())
	fmt.Println("Trying to change card number")
	err := credit.SetCardNumber("invalid")
	if err != nil {
		fmt.Printf("That didn't work: %v\n", err)
	}
	fmt.Printf("Available credit: %v\n", credit.AvailableCredit())
}
