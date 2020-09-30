/*
 * @Author: seekwe
 * @Date:   2019-05-09 16:07:46
 * @Last Modified by:   seekwe
 * @Last Modified time: 2019-05-24 15:08:01
 */

package znet

import (
	"fmt"

	"github.com/sohaha/zlsgo/zlog"
	"github.com/sohaha/zlsgo/zstring"
)

// func debugShowRequest(c *Context) {
// 	method := zstring.Buffer()
// 	method.WriteString("{")
// 	method.WriteString(c.Request.Method)
// 	method.WriteString("}")
// 	c.Log.Debugf("Visit: [%d] %s %s\n", c.Code, zstring.Pad(method.String(), 8, " ", 0), c.Request.RequestURI)
// }

func showRouteDebug(log *zlog.Logger, tf, method, path string) string {
	mLen := zstring.Len(method)
	var mtd string
	min := 6
	if mLen < min {
		mtd = zstring.Pad(method, min, " ", 1)
	} else {
		mtd = zstring.Substr(method, 0, min)
	}

	switch method {
	case "GET":
		method = log.ColorTextWrap(zlog.ColorLightCyan, mtd)
	case "POST":
		method = log.ColorTextWrap(zlog.ColorLightBlue, mtd)
	case "PUT":
		method = log.ColorTextWrap(zlog.ColorLightGreen, mtd)
	case "DELETE":
		method = log.ColorTextWrap(zlog.ColorRed, mtd)
	case "Any":
		method = log.ColorTextWrap(zlog.ColorLightMagenta, mtd)
	case "OPTIONS":
		method = log.ColorTextWrap(zlog.ColorLightMagenta, mtd)
	default:
		method = log.ColorTextWrap(zlog.ColorDefault, mtd)
	}

	return fmt.Sprintf(tf, method, path)
}
