package model

type MovementCard struct {
	ID        interface{}
	Name      interface{}
	Type      interface{}
	Class     interface{}
	Power     interface{}
	Precision interface{}
	Effect    interface{}
	Priority  interface{}
}

var emptyMovement = &MovementCard{}

func (m *MovementCard) ResetStruct() {

	*m = *emptyMovement

}
