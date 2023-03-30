package data_saver

import (
	"context"
	"fmt"
	"os"

	firestore "cloud.google.com/go/firestore"
	model "example.com/AlejandroWaiz/DataManager/Model"
	"google.golang.org/api/option"
)

type AllDataSaverImplementation struct {
	client    *firestore.Client
	ctx       context.Context
	projectID string
}

type AllDataSaver interface {
	SaveAllMovementsIntoFirestore(movements []model.MovementCard) []error
	SaveAllWeathersIntoFirestore(weathers []model.WeatherCard) []error
	SaveAllPokemonsIntoFirestore(pokemons []model.PokemonCard) []error
	SaveAllItemsIntoFirestore(items []model.ItemCard) []error
}

func GetAllDataSaverImplementation() (AllDataSaver, error) {

	ctx := context.Background()

	firestoreProjectID := os.Getenv("firestore_projectID")

	firestoreClient, err := firestore.NewClient(ctx, firestoreProjectID, option.WithCredentialsFile(os.Getenv("serviceAccount_path")))

	if err != nil {
		return nil, fmt.Errorf("[GetAllDataSaverImplementation] error: %v", err)
	}

	return &AllDataSaverImplementation{client: firestoreClient, ctx: ctx, projectID: firestoreProjectID}, nil

}

var isAnError bool = false
