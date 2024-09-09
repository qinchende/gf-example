package logic

import "github.com/qinchende/gofast/fst"

func AddHeaders(c *fst.Context) {
	c.SetHeader("Access-Control-Allow-Origin", "*")
}
