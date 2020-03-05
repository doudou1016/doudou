package ginplus

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		loger := map[string]interface{}{}
		loger["path"] = c.Request.URL.Path
		loger["method"] = c.Request.Method
		loger["ip"] = c.ClientIP()

		loger["url"] = c.Request.URL.String()
		loger["proto"] = c.Request.Proto
		loger["header"] = c.Request.Header
		loger["agent"] = c.Request.UserAgent()

		if m := c.Request.Method; m == http.MethodPost || m == http.MethodPut {
			mediaType, _, _ := mime.ParseMediaType(c.GetHeader("Content-Type"))
			if mediaType == "application/json" {
				body, err := ioutil.ReadAll(c.Request.Body)
				if err == nil {
					c.Request.Body.Close()
					buf := bytes.NewBuffer(body)
					c.Request.Body = ioutil.NopCloser(buf)
					loger["content_length"] = c.Request.ContentLength
					loger["params"] = string(body)
				}
			}
		}

		c.Next()
		loger["status"] = c.Writer.Status()
		loger["length"] = c.Writer.Size()
		loger["latency"] = fmt.Sprintf("%dms", time.Since(start).Nanoseconds()/1e6)
		loger["rtime"] = time.Now().Format("2006-01-02 15:04:05")

		b, err := json.Marshal(loger)
		if err != nil {
			fmt.Println("json.Marshal failed:", err)
			return
		}

		fmt.Println("b:", string(b))

	}
}
