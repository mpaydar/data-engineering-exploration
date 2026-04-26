package main
import(
	   "context"
	   "flag"
	   "fmt"
	   "io"
	   "os"
	   "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)






func main() {
	pr, pw := io.Pipe()
	numPtr := flag.Int("count",10,"number of customers to generate")
	containerPtr := flag.String("container", "lab-data", "target Azure Blob container")
	blobPtr := flag.String("blob", "generated_customers.csv", "target blob name")
	flag.Parse()

	connStr := os.Getenv("AZURE_STORAGE_CONNECTION_STRING")
	if connStr == "" {
		fmt.Println("missing AZURE_STORAGE_CONNECTION_STRING environment variable")
		return
	}

	client, err := azblob.NewClientFromConnectionString(connStr, nil)
	if err != nil {
		fmt.Printf("could not create Azure Blob client: %v\n", err)
		return
	}

	go lab1(*numPtr, pw)

	ctx := context.Background()
	_, err = client.UploadStream(ctx, *containerPtr, *blobPtr, pr, nil)
	if err != nil {
		fmt.Printf("could not upload CSV stream to blob: %v\n", err)
		return
	}

	fmt.Printf("uploaded stream to %s/%s\n", *containerPtr, *blobPtr)
}
