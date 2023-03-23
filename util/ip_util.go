package util

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// Extract Ip Address from HTTP request
// Code below is referencing this link : https://github.com/sbecker/gin-api-demo/blob/07f9a9242f743fc51ae4a046ee58e12627bad571/util/log.go
func GetIpAddress(c *gin.Context) string {
	// first check the X-Forwarded-For header
	requester := c.Request.Header.Get("X-Forwarded-For")
	// if empty, check the Real-IP header
	if len(requester) == 0 {
		requester = c.Request.Header.Get("X-Real-IP")
	}
	// if the requester is still empty, use the hard-coded address from the socket
	if len(requester) == 0 {
		requester = c.Request.RemoteAddr
	}

	// if requester is a comma delimited list, take the first one
	// (this happens when proxied via elastic load balancer then again through nginx)
	if strings.Contains(requester, ",") {
		requester = strings.Split(requester, ",")[0]
	}

	return requester
}
