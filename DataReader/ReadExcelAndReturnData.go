package data_reader

import (
	"fmt"
	"os"
)

func (a *AllDataReaderImplementation) ReadExcelAndReturnData() (interface{}, error) {

	switch a.excelName {

	case os.Getenv("pokemons_excel_filename"):
		pokemons, err := a.PokemonReader.ReadAllPokemonsFromExcelFile()
		if err != nil {
			return nil, fmt.Errorf("[ReadExcelAndReturnData] => %v", err)
		}
		return pokemons, nil

	case os.Getenv("movements_excel_filename"):
		movements, err := a.MovementReader.ReadAllMovementsFromExcelFile()
		if err != nil {
			return nil, fmt.Errorf("[ReadExcelAndReturnData] => %v", err)
		}
		return movements, nil

	case os.Getenv("weathers_excel_filename"):
		weathers, err := a.WeatherCardReader.ReadAllWeathersFromExcelFile()
		if err != nil {
			return nil, fmt.Errorf("[ReadExcelAndReturnData] => %v", err)
		}
		return weathers, nil

	}

	return nil, nil

}
