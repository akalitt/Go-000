package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github/akalitt/go-errors-example/internal/model"
	"github/akalitt/go-errors-example/internal/service"
	"github/akalitt/go-errors-example/pkg/errno"
	"net/http"
	"strconv"
)

type TagDTO struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	State int    `json:"state"`
}

type TagAPI struct {
	TagService service.TagService
}

func NewTagAPI(t service.TagService) TagAPI {
	return TagAPI{TagService: t}
}

func (t *TagAPI) Create(c *gin.Context) {
	tag := model.Tag{}

	err := c.BindJSON(&tag)
	if err != nil {
		errno.Abort(errno.ErrBind, err, c)
		return
	}
	cr, err := t.TagService.Create(tag)

	if err != nil {
		errno.Abort(errno.ErrInsert, err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": TagDTO{
		State: cr.State,
		Name:  cr.Name,
		ID:    cr.ID,
	}})
}

func (b *TagAPI) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("c.Param is not Valid: %s", c.Param("id")))

		errno.Abort(errno.ErrBadRequestParams, err, c)
		return
	}

	tag, err := b.TagService.GetByID(id)

	if err != nil {
		errno.Abort(errno.ErrRecordNotFound, err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": TagDTO{
		State: tag.State,
		Name:  tag.Name,
		ID:    tag.ID,
	}})
}
