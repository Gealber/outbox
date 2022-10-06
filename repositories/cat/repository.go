package cat

import (
	"context"
	"fmt"

	"github.com/Gealber/outbox/repositories/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *repository {
	return &repository{db: db}
}

// Create creates a new cat :).
func (r *repository) Create(ctx context.Context, cat model.Cat) (*model.Cat, error) {
	if err := r.db.Create(&cat).Error; err != nil {
		return nil, err
	}
	fmt.Println(cat.ID)

	return &cat, nil
}

// Update perform update operation to specified cat :).
func (r *repository) Update(ctx context.Context, id string, cat model.Cat) (*model.Cat, error) {
	cat.ID = id
	err := r.db.Transaction(func(tx *gorm.DB) error {
		// update record in cats table.
		result := tx.Model(&model.Cat{}).Clauses(clause.Returning{}).
			Where("id = ?", id).
			Updates(&cat)

		if result.RowsAffected < 1 {
			return gorm.ErrRecordNotFound
		}

		// create event to store.
		event := model.Event{
			Type:         "update",
			ResourceType: "cat",
			ResourceID:   id,
		}

		// write event in events table.
		if err := tx.Model(&model.Event{}).Create(&event).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &cat, nil
}
