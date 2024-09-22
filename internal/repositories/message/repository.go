package message

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/kozld/svetozar/internal/models"
)

type (
	Repository struct {
		db *gorm.DB
	}
)

func New(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) MustMigrate() {
	if err := r.Migrate(); err != nil {
		panic(err)
	}
}

func (r *Repository) Migrate() error {
	return r.db.AutoMigrate(&models.Message{})
}

func (r *Repository) Update(messages ...*models.Message) error {
	res := r.db.Clauses(
		clause.OnConflict{
			UpdateAll: true,
		},
	).Create(&messages)
	return res.Error
}

func (r *Repository) ByStatus(status models.Status) ([]*models.Message, error) {
	messages := make([]*models.Message, 0)
	res := r.db.Where("status = ?", status).Find(&messages)
	return messages, res.Error
}
