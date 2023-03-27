package data_reader

import (
	"errors"
	"os"

	movement_card_reader "example.com/AlejandroWaiz/DataManager/DataReader/MovementCardReader"
	pokemon_card_reader "example.com/AlejandroWaiz/DataManager/DataReader/PokemonCardReader"
	weather_card_reader "example.com/AlejandroWaiz/DataManager/DataReader/WeatherCardReader"
	"github.com/xuri/excelize/v2"
)

type AllDataReaderImplementation struct {
	MovementReader    movement_card_reader.MovementCardReader
	WeatherCardReader weather_card_reader.WeatherCardReader
	PokemonReader     pokemon_card_reader.PokemonCardReader
	excelName         string
}

type AllDataReader interface {
	ReadExcelAndReturnData(excelFile *excelize.File) (interface{}, error)
}

func GetAllDataReaderImplementationToRead(excelName string, excelFile *excelize.File) (AllDataReader, error) {

	var reader AllDataReaderImplementation

	switch excelName {
	case os.Getenv("pokemons_excel_filename"):
		reader.PokemonReader = pokemon_card_reader.GetPokemonCardReaderImplementation(excelFile)
		return &reader, nil

	case os.Getenv("movements_excel_filename"):
		reader.MovementReader = movement_card_reader.GetMovementCardReaderImplementation(excelFile)
		return &reader, nil

	case os.Getenv("weathers_excel_filename"):
		reader.WeatherCardReader = weather_card_reader.GetWeatherCardReaderImplementation(excelFile)
		return &reader, nil

	}

	return nil, errors.New("Not a valid name of excelFile")

}
