package main

import (
	"doudou/router"
	"fmt"

	ginplus "github.com/dllgo/go-gin"
	redisplus "github.com/dllgo/go-redis"
)

func init() {
	//
	redisplus.DefaultClient()

}
func main() {
	//

	mconf := ginplus.Config{Address: ":9090", ReadTimeout: 30, WriteTimeout: 30}
	httpserver, err := ginplus.NewServerHttp(mconf)
	if err != nil {
		fmt.Println(err)
		return
	}
	httpserver.Router = router.InitRouter()
	err = httpserver.Listen()
	defer httpserver.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
