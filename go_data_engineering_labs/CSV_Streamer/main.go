package main
import("fmt"
	 
)






func main() {
	numberOfCustomers := 10
	customers_records := lab1()
	for i:=0 ; i<numberOfCustomers ; i++{
		fmt.Printf("Customer number %d\n", i+1)
		fmt.Println(customers_records[i])
	}
	writeToCSV(customers_records)


}
