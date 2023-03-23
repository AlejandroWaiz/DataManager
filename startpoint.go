package p

import (
	"context"

	"github.com/joho/godotenv"
)

type GCSEvent struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

func StartPoint(ctx context.Context, eventPayload GCSEvent) error {

	godotenv.Load(".env")

	return nil

}
