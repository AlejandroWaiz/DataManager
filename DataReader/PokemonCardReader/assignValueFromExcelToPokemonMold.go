package pokemon_card_reader

import (
	model "example.com/AlejandroWaiz/DataManager/Model"
)

var mapOfPokemonMovementsFromDatabase map[string]string

func assignValueFromDatabaseToPokemonMold(pokemon model.PokemonCard, columnPosition int, columnValue string) {

	if mapOfPokemonMovementsFromDatabase == nil {
		mapOfPokemonMovementsFromDatabase = make(map[string]string)
	}

	if columnPosition == 0 {

		pokemonMold.ID = columnValue

	} else if columnPosition == 1 {

		pokemonMold.Name = columnValue

	} else if columnPosition == 2 {

		pokemonMold.Types = columnValue

	} else if columnPosition == 3 {

		pokemonMold.Pasive = columnValue

	} else if columnPosition == 4 {

		pokemonMold.Stage = columnValue

	} else if columnPosition == 5 {

		pokemonMold.HP = columnValue

	} else if columnPosition == 6 {

		pokemonMold.HpLevelBonus = columnValue

	} else if columnPosition == 7 {

		pokemonMold.Speed = columnValue

	} else if columnPosition == 8 {

		pokemonMold.SpeedLevelBonus = columnValue

	} else if columnPosition == 9 {

		pokemonMold.AttackProficiency = columnValue

	} else if columnPosition == 10 {

		pokemonMold.SpecialAttackProficiency = columnValue

	} else if columnPosition == 11 {

		pokemonMold.DefenseProficiency = columnValue

	} else if columnPosition == 12 {

		pokemonMold.SpecialDefenseProficiency = columnValue

	} else if columnPosition == 13 {

		pokemonMold.SpeedProficiency = columnValue

	} else if columnPosition == 14 {

		mapOfPokemonMovementsFromDatabase["Base movements"] = string(columnValue)

	} else if columnPosition == 15 {

		mapOfPokemonMovementsFromDatabase["Level one movements"] = string(columnValue)

	} else if columnPosition == 16 {

		mapOfPokemonMovementsFromDatabase["Level two movements"] = string(columnValue)

	} else if columnPosition == 17 {

		mapOfPokemonMovementsFromDatabase["Level three movements"] = string(columnValue)

	} else if columnPosition == 18 {

		mapOfPokemonMovementsFromDatabase["Level four movements"] = string(columnValue)

	} else if columnPosition == 19 {

		mapOfPokemonMovementsFromDatabase["Level five movements"] = string(columnValue)

	} else if columnPosition == 20 {

		pokemonMold.Zone = columnValue

	} else if columnPosition == 21 {

		pokemonMold.Rarity = columnValue

	} else if columnPosition == 22 {

		pokemonMold.ApparitionRatio = columnValue
	}

}
