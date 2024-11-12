package main

import "fmt"

type paymenter interface {
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct{
	gateway paymenter
}

// Open close principles
func (p payment) makePayment(amount float32) {
	// razerpayPaymentGw := razorpay{}
	// razerpayPaymentGw.pay(amount)

	// stripePaymentGw := stripe{}
	// stripePaymentGw.pay(amount)

	p.gateway.pay(amount)
}

func (p payment) refundMoney(amount float32, account string) {
	p.gateway.refund(amount,account)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	// logic to make payment
	fmt.Println("making payment using razerpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32) {
	fmt.Println("making payment using stripe", amount)
}

type fakepayment struct{}

func (f fakepayment) pay(amount float32) {
	fmt.Println("making payment using fake gateway for testing purpose", amount)
}

type paypal struct {}

func (p paypal) pay(amount float32) {
	fmt.Println("making payment using paypal", amount)
}

func (p paypal) refund(amount float32, account string) {
	fmt.Println(amount,"refunded to account holder",account)
}

func main() {
	// stripePaymentGw := stripe{}
	// razorpayPayemtGw := razorpay{}
	// fakeGw := fakepayment{}
	paypalPaymentGw := paypal{}
	newPayment := payment{
		gateway: paypalPaymentGw,
	}
	newPayment.makePayment(100)
	newPayment.refundMoney(50,"John")
}