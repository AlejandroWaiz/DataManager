package item_card_reader

import (
	model "example.com/AlejandroWaiz/DataManager/Model"
	"github.com/xuri/excelize/v2"
)

type ItemCardReaderImplementation struct {
	excelFile *excelize.File
}

type ItemCardReader interface {
	ReadAllItemsFromExcelFile() ([]model.ItemCard, error)
}

func GetMovementCardReaderImplementation(file *excelize.File) ItemCardReader {

	return &ItemCardReaderImplementation{excelFile: file}

}
