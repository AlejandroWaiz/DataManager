package model

type Movement struct {
	ID        interface{}
	Name      interface{}
	Type      interface{}
	Class     interface{}
	Power     interface{}
	Precision interface{}
	Effect    interface{}
	Priority  interface{}
}

var emptyMovement = &Movement{}

func (m *Movement) ResetStruct() {

	*m = *emptyMovement

}
