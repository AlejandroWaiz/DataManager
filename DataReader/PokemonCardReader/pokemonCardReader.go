package pokemon_card_reader

import (
	model "example.com/AlejandroWaiz/DataManager/Model"
	"github.com/xuri/excelize/v2"
)

type PokemonCardReaderImplementation struct {
	excelFile *excelize.File
}

type PokemonCardReader interface {
	ReadAllPokemonsFromExcelFile() ([]model.Pokemon, []error)
}

func GetPokemonCardReaderImplementation(file *excelize.File) PokemonCardReader {

	return &PokemonCardReaderImplementation{excelFile: file}

}
