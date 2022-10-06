package msgrelay

import (
	"context"
	"fmt"
	"log"

	"github.com/Gealber/outbox/repositories/model"
)

type eventRepo interface {
	List(ctx context.Context) ([]*model.Event, error)
	Delete(ctx context.Context, ids []string) error
}

// Poll performns POLLING PUBLISHER PATTERN, in a very simple and inefficient manner :).
func Poll(ctx context.Context, eventRepo eventRepo) error {
	log.Println("EXECUTING MSG RELAY...")
	// list events unpublished.
	events, err := eventRepo.List(ctx)
	if err != nil {
		return err
	}

	if len(events) == 0 {
		return nil
	}

	ids := make([]string, 0, len(events))

	// try to publish them.
	for _, event := range events {
		fmt.Printf("PUBLISHING EVENT INTO BROKER: %+v\n", event)
		ids = append(ids, event.ID)
	}	

	// delete them from outbox db.
	return eventRepo.Delete(ctx, ids)
}
