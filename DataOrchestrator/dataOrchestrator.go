package data_orchestrator

import (
	"fmt"

	"cloud.google.com/go/storage"
	data_reader "example.com/AlejandroWaiz/DataManager/DataReader"
	data_saver "example.com/AlejandroWaiz/DataManager/DataSaver"
	"github.com/xuri/excelize/v2"
)

type DataOrchestratorImplementation struct {
	datareader            data_reader.AllDataReader
	datasaver             data_saver.AllDataSaver
	excelFile             *excelize.File
	excelName, bucketName string
}

type DataOrchestrator interface {
	Start() error
}

var localStorageClient *storage.Client

func GetDataOrchestratorImplementation(excelName, bucketName string) (DataOrchestrator, error) {

	orchestrator := DataOrchestratorImplementation{excelName: excelName, bucketName: bucketName}

	excelFile, err := getExcelFile(orchestrator.excelName, orchestrator.bucketName)

	if err != nil {
		return nil, fmt.Errorf("[GetDataOrchestratorImplementation] => %v", err)
	}

	orchestrator.excelFile = excelFile

	reader, err := data_reader.GetAllDataReaderImplementationToRead(orchestrator.excelName, orchestrator.excelFile)

	if err != nil {
		return nil, fmt.Errorf("[GetDataOrchestratorImplementation] => %v", err)
	}

	saver, err := data_saver.GetAllDataSaverImplementation()

	if err != nil {
		return nil, fmt.Errorf("[GetDataOrchestratorImplementation] => %v", err)
	}

	orchestrator.datareader = reader
	orchestrator.datasaver = saver

	return &orchestrator, nil

}
