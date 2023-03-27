package model

type ItemCard struct {
	ID          interface{}
	Name        interface{}
	Effect      interface{}
	Description interface{}
	Cost        interface{}

	//Rareza del item??
}

var emptyItem = &ItemCard{}

func (i *ItemCard) ResetStruct() {

	*i = *emptyItem

}
