package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	updateToScript = `
<script>document.getElementById("msg").innerHTML='%d';</script>
`
	redirect = `<script>self.location='https://gocn.vip/'</script>`
)

func main() {

	engine := gin.Default()

	// HTTP 分块发送实现网页倒计时
	engine.GET("/", func(context *gin.Context) {
		h1 := []byte("<h1 id='msg'></h1>")

		context.Header("Transfer-Encoding", "chunked")
		_, err := context.Writer.Write(h1)
		if err != nil {
			log.Printf("[index] write failed, err: %v", err)
			return
		}

		for i := 1; i <= 5; i++ {
			_, _ = context.Writer.Write([]byte(fmt.Sprintf(updateToScript, i)))
			// Flush sends any buffered data to the client.
			context.Writer.Flush()
			time.Sleep(time.Second)
		}

		_, _ = context.Writer.Write([]byte(redirect))
		context.Writer.Flush()
	})

	err := engine.Run(":8080")
	if err != nil {
		log.Printf("start engine failed, err: %v", err)
		return
	}
}
