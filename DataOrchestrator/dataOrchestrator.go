package data_orchestrator

import (
	data_reader "example.com/AlejandroWaiz/DataManager/DataReader"
	data_saver "example.com/AlejandroWaiz/DataManager/DataSaver"
)

type DataOrchestratorImplementation struct {
	datareader data_reader.AllDataReader
	datasaver  data_saver.AllDataSaver
}

type DataOrchestrator interface {
	ReceiveDataAndSendItToDatabase(data interface{}, excelName string) error
}
