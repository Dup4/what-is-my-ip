package main

import (
	"context"
	"fmt"
	"net"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func handleWhatIsMyIP(c context.Context, ctx *app.RequestContext) {
	ip, err := getIP(c, ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(consts.StatusOK, utils.H{"ip": ip})
}

// https://golangbyexample.com/golang-ip-address-http-request/
func getIP(c context.Context, ctx *app.RequestContext) (string, error) {
	{
		// Get IP from the X-REAL-IP header
		X_REAL_IP := string(ctx.GetHeader("X-REAL-IP"))
		netIP := net.ParseIP(X_REAL_IP)
		if netIP != nil {
			return X_REAL_IP, nil
		}
	}

	{
		// Get IP from X-FORWARDED-FOR header
		ips := string(ctx.GetHeader("X-FORWARDED-FOR"))
		splitIps := strings.Split(ips, ",")
		for _, ip := range splitIps {
			netIP := net.ParseIP(ip)
			if netIP != nil {
				return ip, nil
			}
		}
	}

	{
		// Get IP from RemoteAddr
		ip, _, err := net.SplitHostPort(ctx.RemoteAddr().String())
		if err != nil {
			return "", err
		}

		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	return "", fmt.Errorf("No valid ip found")
}
