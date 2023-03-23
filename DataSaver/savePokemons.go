package data_saver

import (
	"fmt"
	"os"

	model "example.com/AlejandroWaiz/DataManager/Model"
)

func (s *AllDataSaverImplementation) SaveAllPokemonsIntoFirestore(pokemons []model.PokemonCard) []error {

	var arrayOfErrors []error

	for _, pokemon := range pokemons {

		_, err := s.client.Collection(os.Getenv("pokemons_collection_name")).Doc(fmt.Sprint(pokemon.ID)).Set(s.ctx, pokemon)

		if err != nil {
			arrayOfErrors = append(arrayOfErrors, fmt.Errorf("[SaveAllPokemonsIntoFirestore] error saving pokemon %v: %v", pokemon.Name, err))
		}

	}

	if len(arrayOfErrors) > 0 {
		return arrayOfErrors
	}

	return nil

}
