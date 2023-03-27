package data_saver

import (
	"fmt"
	"os"

	model "example.com/AlejandroWaiz/DataManager/Model"
)

func (s *AllDataSaverImplementation) SaveAllItemsIntoFirestore(items []model.ItemCard) []error {

	var arrayOfErrors []error

	for _, item := range items {

		_, err := s.client.Collection(os.Getenv("items_collection_name")).Doc(fmt.Sprint(item.ID)).Set(s.ctx, item)

		if err != nil {
			arrayOfErrors = append(arrayOfErrors, fmt.Errorf("[SaveAllItemsIntoFirestore] error saving item %v: %v", item.Name, err))
		}

	}

	if len(arrayOfErrors) > 0 {
		return arrayOfErrors
	}

	return nil

}
