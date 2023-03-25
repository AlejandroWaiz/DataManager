package data_reader

import "github.com/xuri/excelize/v2"

func (a *AllDataReaderImplementation) ReadExcelAndReturnData(excelFile *excelize.File) (interface{}, error) {

	switch a.excelName {

	case "Pokemon.xlsx":
		a.PokemonReader.ReadAllPokemonsFromExcelFile()

	}

	return nil, nil

}
