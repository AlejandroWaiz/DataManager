package weather_card_reader

import (
	model "example.com/AlejandroWaiz/DataManager/Model"
	"github.com/xuri/excelize/v2"
)

type WeatherCardReaderImplementation struct {
	excelFile *excelize.File
}

type WeatherCardReader interface {
	ReadAllWeathersFromExcelFile() ([]model.WeatherCard, error)
}

func GetWeatherCardReaderImplementation(excelFile *excelize.File) WeatherCardReader {

	return &WeatherCardReaderImplementation{excelFile: excelFile}

}
