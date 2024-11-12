package main

import (
	"fmt"
	"time"
)

// composition
type customer struct {
	name string
	phone string
}

// order struct
type order struct {
	id string
	amount float32
	status string
	createdAt time.Time // nanosecond precision
	customer
}

func newOrder(id string, amount float32, status string) *order {
	// initial setup goes here...
	myOrder := order{
		id: id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}
// Only use pointer when you want to modify struct values
func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	// myOrder := order{
	// 	id: "1",
	// 	amount: 50.00,
	// 	status: "received",
	// }

	// myOrder.createdAt = time.Now()
	// fmt.Println(myOrder.status)

	// myOrder2 := order {
	// 	id: "2",
	// 	amount: 100.00,
	// 	status: "delivered",
	// 	createdAt: time.Now(),
	// }

	// myOrder.status = "paid"

	// fmt.Println("Order struct",myOrder2)
	// fmt.Println("Order struct",myOrder)

	// if you don't set any fields, default value is zero value
	// int => 0, float => 0, string => "", bool => false
	// myOrder := order{
	// 	id: "1",
	// 	amount: 50.00,
	// 	status: "received",
	// }

	// myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder.getAmount())
	// fmt.Println(myOrder)

	// myOrder := newOrder("1",30.50,"received")
	// fmt.Println(myOrder.amount)

	// inline struct
	// language := struct {
	// 	name string
	// 	isGood bool
	// }{"golang",true}
	// fmt.Println(language)

	// newCustomer := customer {
	// 	name: "john",
	// 	phone: "1234567890",
	// }
	myOrder := order {
		id: "1",
		amount: 30,
		status: "received",
		customer: customer {
			name: "john",
			phone: "1234567890",
		},
	}

	myOrder.customer.name = "robin"
	fmt.Println(myOrder.customer)
	fmt.Println(myOrder)
}