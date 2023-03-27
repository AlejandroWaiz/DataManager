package data_saver

import (
	"fmt"
	"os"

	model "example.com/AlejandroWaiz/DataManager/Model"
)

func (s *AllDataSaverImplementation) SaveAllWeathersIntoFirestore(weathers []model.WeatherCard) []error {

	var arrayOfErrors []error

	for _, weather := range weathers {

		_, err := s.client.Collection(os.Getenv("weathers_collection_name")).Doc(fmt.Sprint(weather.ID)).Set(s.ctx, weather)

		if err != nil {
			arrayOfErrors = append(arrayOfErrors, fmt.Errorf("[SaveAllWeathersIntoFirestore] error saving pokemon %v: %v", weather.Name, err))
		}

	}

	if len(arrayOfErrors) > 0 {
		return arrayOfErrors
	}

	return nil

}
