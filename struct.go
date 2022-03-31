package main

import "fmt"

//Insert struct declaration here

func main() {
	customer1 := customer{
		fname:            "Annakin",
		lname:            "Skywalker",
		age:              45,
		subscriber:       true,
		homeAddress:      "Death Star",
		phone:            1234567,
		availCred:        10000,
		currentCartCost:  0.00,
		currentOrderCost: 0.00,
	}

	customer2 := customer{
		fname:            "Han",
		lname:            "Solo",
		age:              50,
		subscriber:       false,
		homeAddress:      "Tatooine",
		phone:            4321765,
		availCred:        20000,
		currentCartCost:  0.00,
		currentOrderCost: 0.00,
	}

	// fmt.Println(customer1.fname, customer2.fname)

	customerlist := []customer{customer1, customer2} // slice of customer struct
	//can append to add more struct customer
	for _, val := range customerlist {
		fmt.Println(val.fname)
	}
}
