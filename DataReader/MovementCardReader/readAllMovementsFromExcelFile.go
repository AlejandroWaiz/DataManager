package movement_card_reader

import (
	"log"
	"os"
	"strings"

	model "example.com/AlejandroWaiz/DataManager/Model"
)

var arrayOfMovements []model.MovementCard
var movementMold model.MovementCard

func (m *MovementCardReaderImplementation) ReadAllMovementsFromExcelFile() ([]model.MovementCard, []error) {

	arrayOfMovementsSheetsFromExcel := strings.Split(os.Getenv("movements_excel_sheetnames"), ",")

	var arrayOfErrors []error

	for _, thisSheet := range arrayOfMovementsSheetsFromExcel {

		allRows, err := m.excelFile.GetRows(thisSheet)

		if err != nil {
			arrayOfErrors = append(arrayOfErrors, err)
			break
		}

		for i := range allRows {

			//Se ignora la primera columna, donde va el nombre de las columnas eg: ID, Name, Type, etc...
			if i == 0 {
				continue
			}

			//Aqui se itera sobre las columnas, tomando el valor de cada una y asignandola a una struct molde, la cual se almacena
			//en un array de pokemones el cual será entregado al adaptador de salida para ser almacenado.
			for _, onlyColumns := range allRows {

				for columnPosition, columnValue := range onlyColumns {

					err := assignValueFromExcelToMovementMold(movementMold, columnPosition, columnValue)

					if err != nil {

						log.Printf("Err assigning value %v from excel to movement: %v ", columnValue, err)

						arrayOfErrors = append(arrayOfErrors, err)

						break

					}

				}

			}

			arrayOfMovements = append(arrayOfMovements, movementMold)

			movementMold.ResetStruct()

		}

	}

	if len(arrayOfErrors) > 0 {
		log.Printf("Errs reading movements: %v", arrayOfErrors)
		return nil, arrayOfErrors
	}

	return arrayOfMovements, nil

}
