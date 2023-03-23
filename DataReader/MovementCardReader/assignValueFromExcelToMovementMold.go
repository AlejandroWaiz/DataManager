package movement_card_reader

import (
	model "example.com/AlejandroWaiz/DataManager/Model"
)

func assignValueFromExcelToMovementMold(movementMold model.Movement, columnPosition int, columnValue string) error {

	//Actually liked more this assignation than pokemon's one
	switch columnPosition {

	case 0:
		movementMold.ID = columnValue

	case 1:
		movementMold.Name = columnValue

	case 2:
		movementMold.Type = columnValue

	case 3:
		movementMold.Class = columnValue

	case 4:
		movementMold.Power = columnValue

	case 5:
		movementMold.Precision = columnValue

	case 6:
		movementMold.Priority = columnValue

	case 7:
		movementMold.Effect = columnValue
	}

	return nil
}
