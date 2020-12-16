package errno

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Error struct {
	Code       string `json:"code"`
	StatusCode int    `json:"status_code"`
	MessageEN  string `json:"message_en"`
	MessageCN  string `json:"message_cn"`
}

// Err represents an error, `Code`, `File`, `Line`, `Func` will be automatically filled.
type Err struct {
	StatusCode int    `json:"status_code"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Detail     string `json:"detail"`
	File       string `json:"file"`
}

// Error returns the error message.
func (e *Err) Error() string {
	return e.Message
}

// Fill the error struct with the detail error information.

// Abort the current request with the specified error code.
func Abort(code string, err error, c *gin.Context) {
	e := &Err{Code: code}
	d := ERROR_MESSAGE[code]

	e.StatusCode = d.StatusCode
	e.Code = d.Code
	e.Message = d.MessageCN

	fmt.Printf("Abort %+v\n", err)
	c.Error(err)

	c.AbortWithStatusJSON(e.StatusCode, e)
}
