package main
import (  "encoding/csv"
		  "os"
          "strconv"
		  "strings"
		)


func writeToCSV(customers []customers) error {
    // 1. Create the file
    file, err := os.Create("customers.csv")
    if err != nil {
        return err
    }
    defer file.Close() // Best practice: ensure file closes when done

    // 2. Create the CSV writer
    writer := csv.NewWriter(file)
    defer writer.Flush() // Push any leftover data in the buffer to the file

    // 3. Write the Header Row
    headers := []string{"First Name", "Last Name", "Purchased", "Items"}
    writer.Write(headers)

    // 4. Loop through your list and write rows
    for _, c := range customers {
        row := []string{
            c.first_name,
            c.last_name,
            strconv.FormatBool(c.purchased),
			strings.Join(c.items_bought, ","),
            strconv.Itoa(len(c.items_bought)), // Convert length to string
        }
        writer.Write(row)
    }

    return nil
}