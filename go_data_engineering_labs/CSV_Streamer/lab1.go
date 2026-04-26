package main

import (
	"github.com/brianvoe/gofakeit/v7";
	// "regexp"
	"math/rand/v2"
	"io"
	"encoding/csv"
	"strconv"
	"strings"
)
type customers struct {
	first_name string
	last_name string 
	purchased bool
	items_bought []string
	// item_category string
}
func lab1(numCustomers int,pw *io.PipeWriter){
	csvWriter := csv.NewWriter(pw)
	for i:=0;i<numCustomers;i++ {
		newCustomer:=customers{}
		number_of_items := generate_random_number()
		newCustomer.first_name=gofakeit.FirstName()
		newCustomer.last_name=gofakeit.LastName()
		newCustomer.purchased=gofakeit.Bool()
		customer_cart := pick_item(number_of_items)
		newCustomer.items_bought=customer_cart
		row := []string{newCustomer.first_name, newCustomer.last_name,strconv.FormatBool(newCustomer.purchased),strings.Join(newCustomer.items_bought,","),strconv.Itoa(number_of_items)}
		csvWriter.Write(row)
		}
	csvWriter.Flush() // Ensure the last bits of text leave the buffer
	pw.Close()        // Tell the other end "No more data is coming!"
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