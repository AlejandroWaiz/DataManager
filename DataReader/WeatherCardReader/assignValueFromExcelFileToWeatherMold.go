package weather_card_reader

import model "example.com/AlejandroWaiz/DataManager/Model"

func assignValueFromExcelToWeatherMold(weatherMold model.WeatherCard, columnPosition int, columnValue string) {

	switch columnPosition {

	case 0:
		weatherMold.ID = columnValue

	case 1:
		weatherMold.Name = columnValue

	case 2:
		weatherMold.Duration = columnValue

	case 3:
		weatherMold.Effect = columnValue

	case 4:
		weatherMold.Description = columnValue

	}

}
