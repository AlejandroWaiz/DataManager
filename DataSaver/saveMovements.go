package data_saver

import (
	"fmt"
	"os"

	model "example.com/AlejandroWaiz/DataManager/Model"
)

func (s *AllDataSaverImplementation) SaveAllMovementsIntoFirestore(movements []model.MovementCard) []error {

	var arrayOfErrors []error

	for _, movement := range movements {

		_, err := s.client.Collection(os.Getenv("movements_collection_name")).Doc(fmt.Sprint(movement.ID)).Set(s.ctx, movement)

		if err != nil {
			arrayOfErrors = append(arrayOfErrors, fmt.Errorf("[SaveAllMovementsIntoFirestore] error saving pokemon %v: %v", movement.Name, err))
		}

	}

	if len(arrayOfErrors) > 0 {
		return arrayOfErrors
	}

	return nil

}
