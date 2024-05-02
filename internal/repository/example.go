package repository

import (
	"context"
	"sync"

	"github.com/foo/bar/internal/model"
	"gorm.io/gorm"
)

type ExampleRepository struct {
	mutex sync.Mutex
	db    *gorm.DB
}

func NewExampleRepository(db *gorm.DB) *ExampleRepository {
	return &ExampleRepository{
		db: db,
	}
}

func (r *ExampleRepository) Find(ctx context.Context, id int) (*model.ExampleModel, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	var exampleModel model.ExampleModel

	res := r.db.WithContext(ctx).Take(&exampleModel, id)
	if res.Error != nil {
		return nil, res.Error
	}

	return &exampleModel, nil
}

func (r *ExampleRepository) Create(ctx context.Context, exampleModel *model.ExampleModel) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	res := r.db.WithContext(ctx).Create(exampleModel)

	return res.Error
}
