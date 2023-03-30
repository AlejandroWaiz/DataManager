package pokemon_card_reader

import (
	"fmt"
	"os"
	"strings"

	model "example.com/AlejandroWaiz/DataManager/Model"
)

var arrayOfPokemons []model.PokemonCard
var pokemonMold model.PokemonCard

func (p *PokemonCardReaderImplementation) ReadAllPokemonsFromExcelFile() ([]model.PokemonCard, error) {

	allPokemonExcelSheets := os.Getenv("pokemons_excel_sheetnames")

	arrayOfPokemonSheetsFromExcel := strings.Split(allPokemonExcelSheets, ",")

	//Por cada hoja del excel de Pokemons se hará una iteración, recorriendo el documento en su totalidad.
	for _, thisSheet := range arrayOfPokemonSheetsFromExcel {

		allRows, err := p.excelFile.GetRows(thisSheet)

		if err != nil {
			return nil, fmt.Errorf("[ReadAllPokemonsFromExcelFile] error: %v", err)
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

					assignValueFromDatabaseToPokemonMold(pokemonMold, columnPosition, columnValue)

				}

			}

			pokemonMold.Movements = mapOfPokemonMovementsFromDatabase

			arrayOfPokemons = append(arrayOfPokemons, pokemonMold)
			pokemonMold.ResetStruct()

		}

	}

	return arrayOfPokemons, nil
}
