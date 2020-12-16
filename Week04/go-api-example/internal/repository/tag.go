package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github/akalitt/go-errors-example/internal/model"
)

type TagRepository interface {
	Create(tag model.Tag) (model.Tag, error)
	//GetAll() []model.Tag
	GetByID(id int) (model.Tag, error)
	//Delete(book model.Tag)
	//Update(id int, newTag model.Tag)
}

type tagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return &tagRepository{db: db}
}

func (t *tagRepository) Create(tag model.Tag) (model.Tag, error) {
	if err := t.db.Create(&tag).Error
		err != nil {
		return model.Tag{}, errors.Wrapf(err, "Tag  / Create err: %d: %+v", tag)
	}
	return tag, nil
}

func (t *tagRepository) GetByID(id int) (model.Tag, error) {
	var tag model.Tag

	if err := t.db.Where("id = ? AND is_del = ? ", id, 0).
		Find(&tag).Error; err != nil {
		return model.Tag{}, errors.Wrapf(err, "Tag  / GetByID err: %d: %+v", id, t)
	}
	return tag, nil
}
