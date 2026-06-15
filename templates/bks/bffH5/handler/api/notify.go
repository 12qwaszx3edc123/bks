package api

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func Notify(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		fmt.Println("解析失败:", err)
		c.String(200, "fail")
		return
	}
	notifyMap := make(map[string]string)
	for s, v := range c.Request.PostForm {
		notifyMap[s] = v[0]
	}
	fmt.Println(notifyMap)
	drugSn := notifyMap["out_trade_no"]
	if drugSn == "" {
		fmt.Println("订单号为空")
		c.String(200, "fail")
		return
	}
	if notifyMap["trade_status"] != "TRADE_SUCCESS" {
		fmt.Println("支付失败,trade_status:", notifyMap["trade_status"])
		c.String(200, "fail")
		return
	}
	if strings.Contains(drugSn, "do") {
		fmt.Println("处理药品订单支付回调")
		c.String(200, "success")
		return
	}
}

func HandlerExpireOrder() {
	fmt.Println("暂无失效订单")
}
