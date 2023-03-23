package event_card_reader

type EventCardReaderImplementation struct {
}

type EventCardReader interface{}

func GetEventCardReaderImplementation() EventCardReader {

	return &EventCardReaderImplementation{}

}
