package model

type Pokemon struct {
	ID              interface{}
	Name            interface{}
	Types           interface{}
	Pasive          interface{}
	Stage           interface{}
	HP              interface{}
	HpLevelBonus    interface{}
	Speed           interface{}
	SpeedLevelBonus interface{}
	Zone            interface{}
	Rarity          interface{}
	//From 1 to 100, that will represent the % from his rarity group and zone being:
	ApparitionRatio interface{}
	Movements       []map[string]string
	Proficiencies
}

type Proficiencies struct {
	AttackProficiency, SpecialAttackProficiency, DefenseProficiency, SpecialDefenseProficiency, SpeedProficiency string
}

var emptyPokemon = &Pokemon{}

func (p *Pokemon) ResetStruct() {

	*p = *emptyPokemon

}
