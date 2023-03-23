package p

import "context"

type GCSEvent struct {
	Bucket string `json:"bucket"`
	Name   string `json:"name"`
}

func StartPoint(ctx context.Context, eventPayload GCSEvent) error {

	return nil

}
