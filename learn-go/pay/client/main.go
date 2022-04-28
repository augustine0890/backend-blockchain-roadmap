package main

import (
	"fmt"
	"pay/paybroker"
	"pay/payment"
)

type PaymentOption interface {
	ProcessPayment(float32) bool
}

type Account struct{}

func (c *Account) AvailableFunds() float32 {
	fmt.Println("Listing available funds...")
	return 50
}

func (a *Account) ProcessPayment(amount float32) bool {
	fmt.Println("Processing payment...")
	return true
}

type CreditAccount struct {
	Account
}

type CheckingAccount struct{}

func (c *CheckingAccount) AvailableFunds() float32 {
	fmt.Println("Getting checking funds...")
	return 250
}

type HybirdAccount struct {
	Account
	CheckingAccount
}

func (h *HybirdAccount) AvailableFunds() float32 {
	return h.Account.AvailableFunds() + h.CheckingAccount.AvailableFunds()
}

func (c *CreditAccount) processPayment(amount float32) {
	fmt.Println("Processing credit card payment...")
}

func NewCreditAccount(chargeCh chan float32) *CreditAccount {
	creditAccount := &CreditAccount{}
	go func(chargeCh chan float32) {
		for amount := range chargeCh {
			creditAccount.processPayment(amount)
		}
	}(chargeCh)
	return creditAccount
}

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

	var option PaymentOption

	option = payment.CreateCreditAccount(
		"Debra Ingram",
		"1111-2222-3333-4444",
		5,
		2025,
		123,
	)
	option.ProcessPayment(500)

	option = payment.CreateCashAccount()

	option.ProcessPayment(500)

	option = &paybroker.PaymentBrokerAccount{}

	option.ProcessPayment(500)

	chargeCh := make(chan float32)
	NewCreditAccount(chargeCh)
	chargeCh <- 300
	var a string
	fmt.Scanln(&a)

	ca := &CreditAccount{}
	ca.AvailableFunds()
	ca.ProcessPayment(500)

	ha := &HybirdAccount{}
	fmt.Println(ha.AvailableFunds())
	fmt.Println(ha.CheckingAccount.AvailableFunds())
}
