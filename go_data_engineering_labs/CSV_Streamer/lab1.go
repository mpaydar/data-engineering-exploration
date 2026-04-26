package main

import (
	"github.com/brianvoe/gofakeit/v7";
	// "regexp"
	"math/rand/v2"
)
type customers struct {
	first_name string
	last_name string 
	purchased bool
	items_bought []string
	// item_category string
}
func lab1()  [] customers {
	numberOfCustomers := 10
	customer_transactions := make([]customers, 0, numberOfCustomers)
	for i:=0;i<numberOfCustomers;i++ {
		newCustomer:=customers{}
		number_of_items := generate_random_number()
		newCustomer.first_name=gofakeit.FirstName()
		newCustomer.last_name=gofakeit.LastName()
		newCustomer.purchased=gofakeit.Bool()
		customer_cart := pick_item(number_of_items)
		newCustomer.items_bought=customer_cart
		customer_transactions = append(customer_transactions, newCustomer)	
		}
	

	return customer_transactions
	}



func generate_random_number () int{
	number_of_item_bought := rand.N(20)
	return number_of_item_bought
}

func pick_item(num_item_buying int) [] string {
	items_available := [...] string {"egg","chicken","orange","steak","milk","yogurt","shrimp","parsley","bread"}
	cart := []string{}
	maxItemPurchaseItem := 10
	for i:=0 ; i<num_item_buying && i < maxItemPurchaseItem ; i++ {
		index := rand.N(len(items_available))
		particular_item_bought := items_available[index]
		cart = append(cart, particular_item_bought)
	}
	return cart
}