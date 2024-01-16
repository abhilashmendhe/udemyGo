package main

import "fmt"

/* O - Open-Closed Principle (OCP)
This principle states that a struct should be open for extension but closed for
modification, meaning that the behavior of a struct can be extended without
changing its code. This helps to keep the code flexible and adaptable to changing
requirements.
*/

type PaymentMethod interface {
	Pay()
}

type Payment struct{}

func (p Payment) Process(pm PaymentMethod) {
	pm.Pay()
}

// Credit Card
type CreditCard struct {
	amount float64
}

func (cc CreditCard) Pay() {
	fmt.Printf("Paid %.2f using CreditCard.\n", cc.amount)
}

// Debit Card
type DebitCard struct {
	amount float64
}

func (dc DebitCard) Pay() {
	fmt.Printf("Paid %.2f using DebigCard.\n", dc.amount)
}

/*
As per OCP, my Payment struct is open for extension and closed for modification.
Since I’m using PaymentMethod interface, I don’t have to edit Payment behavior
when adding new payment methods
*/

func main() {
	p := Payment{}
	cc := CreditCard{12.23}
	p.Process(cc)
}
