package weather_card_reader

import (
	"fmt"
	"os"
	"strings"

	model "example.com/AlejandroWaiz/DataManager/Model"
)

var arrayOfWeathers []model.WeatherCard
var weatherMold model.WeatherCard

func (w *WeatherCardReaderImplementation) ReadAllWeathersFromExcelFile() ([]model.WeatherCard, error) {

	arrayOfWeathersSheetsFromExcel := strings.Split(os.Getenv("movements_excel_sheetnames"), ",")

	for _, thisSheet := range arrayOfWeathersSheetsFromExcel {

		allRows, err := w.excelFile.GetRows(thisSheet)

		if err != nil {

			return nil, fmt.Errorf("[ReadAllWeathersFromExcelFile] error: %v", err)

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

					assignValueFromExcelToWeatherMold(weatherMold, columnPosition, columnValue)

				}

			}

			arrayOfWeathers = append(arrayOfWeathers, weatherMold)

			weatherMold.ResetStruct()

		}

	}

	return arrayOfWeathers, nil
}
