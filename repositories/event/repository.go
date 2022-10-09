package events

import (
	"context"

	"github.com/Gealber/outbox/msgrelay"
	"github.com/Gealber/outbox/repositories/model"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repository {
	return &repository{db: db}
}

// List list events according to publish requirement.
func (r *repository) List(ctx context.Context) ([]*model.Event, error) {
	var events []*model.Event

	err := r.db.Model(&model.Event{}).
		Find(&events, "published = ?", false).Error
	if err != nil {
		return nil, err
	}

	return events, nil
}

// Delete events that are already published.
func (r *repository) Delete(ctx context.Context, ids []string) error {
	return r.db.Where("id IN ?", ids).Delete(&model.Event{}).Error
}

// ChangeFeed fetch logs from events feed.
func (r *repository) ChangeFeed(ctx context.Context, pubsub msgrelay.PubSub) error {
	rows, err := r.db.Raw("EXPERIMENTAL CHANGEFEED FOR events").Rows()
	if err != nil {
		return err
	}

	defer rows.Close()

	topic := pubsub.Topic("cats")

	var (
		table string
		key   string
		value []byte
	)

	for rows.Next() {
		if err := rows.Scan(&table, &key, &value); err != nil {
			return err
		}

		pubsub.Publish(ctx, topic, value)
	}

	return nil
}
