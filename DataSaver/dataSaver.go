package data_saver

import (
	"context"

	firestore "cloud.google.com/go/firestore"
)

type AllDataSaverImplementation struct {
	client    *firestore.Client
	ctx       context.Context
	projectID string
}

type AllDataSaver interface {
}
