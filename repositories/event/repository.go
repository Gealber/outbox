package eve_test

import (
	"context"

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