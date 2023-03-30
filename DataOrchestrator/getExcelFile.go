package data_orchestrator

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/storage"
	"github.com/xuri/excelize/v2"
	"google.golang.org/api/option"
)

func getExcelFile(excelName, bucketName string) (*excelize.File, error) {

	ctx := context.Background()

	storageClient, err := storage.NewClient(ctx, option.WithCredentialsFile(os.Getenv("serviceAccount_path")))

	if err != nil {

		return nil, fmt.Errorf("[Cloud Storage|NewClient] err : %v", err)

	}

	localStorageClient = storageClient

	objectReader, err := storageClient.Bucket(bucketName).Object(excelName).NewReader(ctx)

	if err != nil {

		return nil, fmt.Errorf("[Cloud Storage|NewReader] err : %v", err)

	}

	defer objectReader.Close()

	f, err := excelize.OpenReader(objectReader)

	//f, err := excelize.OpenFile("excel-test-4structs with err on numero factura.xlsx")

	if err != nil {

		return nil, fmt.Errorf("[Excelize|OpenReader] err : %v", err)

	}

	return f, nil

}
