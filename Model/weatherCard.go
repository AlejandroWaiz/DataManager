package model

type WeatherCard struct {
	ID, Name, Duration, Description, Effect interface{}
}

var emptyWeatherCard = &WeatherCard{}

func (w *WeatherCard) ResetStruct() {

	*w = *emptyWeatherCard

}
