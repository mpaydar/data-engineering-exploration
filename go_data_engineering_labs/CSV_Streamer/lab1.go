package main

import (
	"github.com/brianvoe/gofakeit/v7";
	"fmt"
)
type customers struct {
	first_name string
	last_name string 
	purchased bool
	items_bought []string
	item_category string
}
func lab1() {
	for i:=0;i<10;i++ {
		first_name := gofakeit.FirstName()
		fmt.Println(first_name)

	}

}

