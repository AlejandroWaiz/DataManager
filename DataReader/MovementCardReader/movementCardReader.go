package movement_card_reader

import (
	model "example.com/AlejandroWaiz/DataManager/Model"
	"github.com/xuri/excelize/v2"
)

type MovementCardReaderImplementation struct {
	excelFile *excelize.File
}

type MovementCardReader interface {
	ReadAllMovementsFromExcelFile() ([]model.MovementCard, []error)
}

func GetMovementCardReaderImplementation(file *excelize.File) MovementCardReader {

	return &MovementCardReaderImplementation{excelFile: file}

}
