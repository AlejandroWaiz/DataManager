package data_saver

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"github.com/xuri/excelize/v2"
)

func storeExcelIntoBucket(excelFile *excelize.File, excelFileName string) error {

	ctx := context.Background()

	storageClient, err := storage.NewClient(ctx)

	if err != nil {

		return fmt.Errorf("[StorExcelIntoBucket] err: %v", err)

	}

	var destinyBucket string

	if isAnError {
		destinyBucket = os.Getenv("bucketForErrors")
	} else {
		destinyBucket = os.Getenv("bucketForExcelOfValidData")
	}

	w, err := excelFile.WriteToBuffer()

	if err != nil {
		log.Printf("Err getting buffer of excelfile: %v", err)
		return err
	}

	o := storageClient.Bucket(destinyBucket).Object(excelFileName)

	wc := o.NewWriter(context.Background())

	wc.ContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"

	log.Println("starting to copy")

	if _, err := io.Copy(wc, w); err != nil {
		log.Printf("Err coping into bucket: %v ", err)
		return fmt.Errorf("io.Copy: %v", err)
	}

	log.Println("COPIED")

	if err := wc.Close(); err != nil {
		log.Printf("Err closing object: %v ", err)
		return fmt.Errorf("Writer.Close: %v", err)
	}

	log.Println("CLOSED")

	return nil

}
