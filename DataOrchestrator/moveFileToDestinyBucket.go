package data_orchestrator

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
)

func (o *DataOrchestratorImplementation) storeExcelIntoBucket() error {

	ctx := context.Background()

	var destinyBucket string

	if isAnErrorHappen {
		destinyBucket = os.Getenv("bucketForErrors")
	} else {
		destinyBucket = os.Getenv("bucketForExcelOfValidData")
	}

	w, err := o.excelFile.WriteToBuffer()

	if err != nil {
		log.Printf("Err getting buffer of excelfile: %v", err)
		return err
	}

	object := localStorageClient.Bucket(destinyBucket).Object(o.excelName)

	wc := object.NewWriter(ctx)

	wc.ContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"

	if _, err := io.Copy(wc, w); err != nil {
		log.Printf("Err coping into bucket: %v ", err)
		return fmt.Errorf("io.Copy: %v", err)
	}

	if err := wc.Close(); err != nil {
		log.Printf("Err closing object: %v ", err)
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil

}
