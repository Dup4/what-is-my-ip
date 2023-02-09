package main

import (
	"context"
	"fmt"

	"github.com/Dup4/what-is-my-ip/version"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func main() {
	fmt.Println(version.ToString())

	h := server.Default()

	h.GET("/", func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	h.GET("/what-is-my-ip", handleWhatIsMyIP)

	h.Spin()
}
