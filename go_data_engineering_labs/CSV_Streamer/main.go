package main
import("fmt"
	   "flag"
	 
)






func main() {
	numPtr := flag.Int("count",10,"number of customers to generate")
	flag.Parse()
	numberOfCustomers := *numPtr
	customers_records := lab1(numberOfCustomers)
	for i:=0 ; i<numberOfCustomers ; i++{
		fmt.Printf("Customer number %d\n", i+1)
		fmt.Println(customers_records[i])
	}
	writeToCSV(customers_records)


}
