//+build

package router

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github/akalitt/go-errors-example/api"
	"github/akalitt/go-errors-example/internal/repository"
	"github/akalitt/go-errors-example/internal/service"
)

func initBookAPI(db *gorm.DB) api.TagAPI {
	wire.Build(repository.NewTagRepository, service.NewTagService, api.NewTagAPI)
	return api.TagAPI{}
}
