package main

import (
	"fmt"
	"log"

	// Imports the Google Cloud BigQuery client package.
	"cloud.google.com/go/bigquery"
	"golang.org/x/net/context"
)

func main() {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "parkingnavigator-92124"

	// Creates a client.
	client, err := bigquery.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name for the new dataset.
	datasetName := "IUPUI2"

	// Creates a Dataset instance.
	dataset := client.Dataset(datasetName)
	schema1 := bigquery.Schema{
		&bigquery.FieldSchema{Name: "Name", Required: true, Type: bigquery.StringFieldType},
		&bigquery.FieldSchema{Name: "score", Repeated: true, Type: bigquery.IntegerFieldType},
		&bigquery.FieldSchema{Name: "Appt", Required: false, Type: bigquery.StringFieldType},
	}
	if err := dataset.Create(ctx); err != nil {
		log.Fatalf("Failed to create dataset: %v", err)
	}
	fmt.Println("Dataset created!")

	// Creates the new BigQuery dataset.
	err = dataset.Table("Test2").Create(ctx, schema1)
	if err != nil {
		log.Fatalf("failed %v", err)
	}

	fmt.Printf("Table created\n")
}
