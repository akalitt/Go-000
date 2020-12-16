package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/willf/pad"
	"github/akalitt/go-errors-example/pkg/errno"
	"github/akalitt/go-errors-example/pkg/logger"
	"io/ioutil"

	"time"
)

func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {

		start := time.Now().UTC()

		path := c.Request.URL.Path
		// Continue.
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) //We have to create a new Buffer, because rdr1 will be read.

		c.Request.Body = rdr2

		bff := new(bytes.Buffer)
		bff.ReadFrom(rdr1)

		s := bff.String()

		c.Next()
		// Skip for the health check requests.
		//if path == "/metrics" || path == "/sd/health" || path == "/sd/ram" || path == "/sd/cpu" || path == "/sd/disk" {
		//	return
		//}
		// Calculates the latency.
		end := time.Now().UTC()
		latency := end.Sub(start)

		// The basic informations.
		status := c.Writer.Status()
		method := c.Request.Method
		ip := c.ClientIP()
		userAgent := c.Request.UserAgent()

		// Create the symbols for each status.
		statusString := ""
		switch {
		case status >= 500:
			statusString = fmt.Sprintf("▲ %d", status)
		case status >= 400:
			statusString = fmt.Sprintf("▲ %d", status)
		case status >= 300:
			statusString = fmt.Sprintf("■ %d", status)
		case status >= 100:
			statusString = fmt.Sprintf("● %d", status)
		}

		// Data fields that will be recorded into the log files.

		fields := logrus.Fields{
			"user_agent":   userAgent,
			"request_body": fmt.Sprintf("%s", s),
		}

		fmt.Printf("c.Erroes \n", len(c.Errors))
		fmt.Printf("c.Erro LEN == 0 \n", len(c.Errors) == 0)

		// Append the error to the fields so we can record it.
		if len(c.Errors) != 0 {
			for k, v := range c.Errors {
				// Skip if it's the Gin internal error.
				if !v.IsType(gin.ErrorTypePrivate) {
					continue
				}
				// The field name with the `error_INDEX` format.
				errorKey := fmt.Sprintf("error_%d", k)

				switch v.Err.(type) {
				case *errno.Err:
					e := v.Err.(*errno.Err)
					fields[errorKey] = fmt.Sprintf("%s[%s:%d:%s]", e.Code, e.File)
				default:
					fields[errorKey] = fmt.Sprintf("%+v", v.Err)
				}
			}
		}

		msg := fmt.Sprintf("%s | %13s | %12s | %s %s ", statusString, latency, ip, pad.Right(method, 2, " "), path)
		if len(c.Errors) == 0 {
			// Example: ● 200 |  102.268592ms |    127.0.0.1 | POST  /user (user_agent=xxx)

			fmt.Println("== 0")
			logger.Infof(msg, fields)
		} else {
			fmt.Println("!= 0 logger.Errorf")

			// Example: ▲ 403 |  102.268592ms |    127.0.0.1 | POST  /user (user_agent=xxx error_0=xxx)
			logger.Errorf(msg, fields)
		}
	}
}
