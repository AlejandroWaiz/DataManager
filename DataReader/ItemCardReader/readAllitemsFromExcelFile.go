package item_card_reader

import (
	"fmt"
	"os"
	"strings"

	model "example.com/AlejandroWaiz/DataManager/Model"
)

var arrayOfItems []model.ItemCard
var itemMold model.ItemCard

func (i *ItemCardReaderImplementation) ReadAllItemsFromExcelFile() ([]model.ItemCard, error) {

	allPokemonExcelSheets := os.Getenv("items_excel_sheetnames")

	arrayOfPokemonSheetsFromExcel := strings.Split(allPokemonExcelSheets, ",")

	for _, thisSheet := range arrayOfPokemonSheetsFromExcel {

		allRows, err := i.excelFile.GetRows(thisSheet)

		if err != nil {
			return nil, fmt.Errorf("[ReadAllItemsFromExcelFile] error: %v", err)
		}

		for i := range allRows {

			//Se ignora la primera columna, donde va el nombre de las columnas eg: ID, Name, Type, etc...
			if i == 0 {
				continue
			}

			//Aqui se itera sobre las columnas, tomando el valor de cada una y asignandola a una struct molde, la cual se almacena
			//en un array de Items el cual será entregado al adaptador de salida para ser almacenado.
			for _, onlyColumns := range allRows {

				for columnPosition, columnValue := range onlyColumns {

					assignValueFromDatabaseToPokemonMold(itemMold, columnPosition, columnValue)

				}

			}

			arrayOfItems = append(arrayOfItems, itemMold)
			itemMold.ResetStruct()

		}

	}

	return arrayOfItems, nil

}
