package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob/bloberror"
	"github.com/joho/godotenv"
)






func main() {
	pr, pw := io.Pipe()
	numPtr := flag.Int("count",10,"number of customers to generate")
	containerPtr := flag.String("container", "lab-data", "target Azure Blob container")
	blobPtr := flag.String("blob", "generated_customers.csv", "target blob name")
	createContainerPtr := flag.Bool("create-container", true, "create the target container if it does not exist")
	flag.Parse()

	// Load local secrets file if present (does not override existing env vars).
	if err := godotenv.Load(".env.local"); err != nil && !errors.Is(err, os.ErrNotExist) {
		fmt.Printf("note: could not load .env.local: %v\n", err)
	}

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

	ctx := context.Background()
	if *createContainerPtr {
		_, err = client.CreateContainer(ctx, *containerPtr, nil)
		if err != nil && !bloberror.HasCode(err, bloberror.ContainerAlreadyExists) {
			fmt.Printf("could not create container %q: %v\n", *containerPtr, err)
			return
		}
	}

	go lab1(*numPtr, pw)

	_, err = client.UploadStream(ctx, *containerPtr, *blobPtr, pr, nil)
	if err != nil {
		fmt.Printf("could not upload CSV stream to blob: %v\n", err)
		return
	}

	fmt.Printf("uploaded stream to %s/%s\n", *containerPtr, *blobPtr)
}
