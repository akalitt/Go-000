package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strconv"
)

// 所有的错误都在controller处理
// 底层只负责把错误传递上来
func GetUserByIdController(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		err = errors.Wrap(err, "invalid userId")
		log.Printf("%+v", err)

		c.AbortWithStatus(http.StatusInternalServerError)
	}

	s := UserService{
		ID: id,
	}

	name, err := s.GetUserByIdService()

	if err != nil {
		log.Printf("%+v", err)
		// 没找到记录 返回404
		if errors.Is(err, sql.ErrNoRows) {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			// 其他错误 直接500
			c.String(http.StatusOK, "something went wrong")
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, name)

}
