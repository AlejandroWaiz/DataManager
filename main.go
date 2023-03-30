package main

import (
	"log"

	dataOrchestrator "example.com/AlejandroWaiz/DataManager/DataOrchestrator"
	"github.com/joho/godotenv"
)

type GCSEvent struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

func main() {

	godotenv.Load(".env")

	orchestrator, err := dataOrchestrator.GetDataOrchestratorImplementation("Pokemons.xlsx", "save-pokemons-here")

	if err != nil {
		log.Println(err)
		panic(err)
	}

	err = orchestrator.Start()

	if err != nil {
		log.Println(err)
		panic(err)
	}

}

/* func StartPoint(ctx context.Context, eventPayload GCSEvent) error {

	godotenv.Load(".env")

	return nil

}
*/
