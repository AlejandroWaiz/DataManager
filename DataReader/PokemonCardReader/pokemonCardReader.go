package pokemon_card_reader

import "github.com/xuri/excelize/v2"

type PokemonCardReaderImplementation struct {
	pokemonExcelFile *excelize.File
}

type PokemonCardReader interface{}
