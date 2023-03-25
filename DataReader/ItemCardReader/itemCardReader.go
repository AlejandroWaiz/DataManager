package item_card_reader

import (
	"github.com/xuri/excelize/v2"
)

type ItemCardReaderImplementation struct {
	excelFile *excelize.File
}

type ItemCardReader interface {
}

func GetMovementCardReaderImplementation(file *excelize.File) ItemCardReader {

	return &ItemCardReaderImplementation{excelFile: file}

}
