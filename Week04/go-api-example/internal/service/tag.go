package service

import (
	"github/akalitt/go-errors-example/internal/model"
	"github/akalitt/go-errors-example/internal/repository"
	//"github/akalitt/go-errors-example/internal/model"
)

type TagService struct {
	TagRepository repository.TagRepository
}

func NewTagService(t repository.TagRepository) TagService {
	return TagService{TagRepository: t}
}

func (t *TagService) Create(tag model.Tag) (model.Tag, error) {
	return t.TagRepository.Create(tag)
}

func (t *TagService) GetByID(id int) (model.Tag, error) {
	return t.TagRepository.GetByID(id)
}
