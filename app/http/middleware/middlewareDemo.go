package middleware

import (
	"fmt"

	"github.com/tianzhaocn/skyscraper/framework"
)

func Mydemo() framework.ControllerHandler {
	// 使用函数回调
	return func(c *framework.Context) error {
		fmt.Println("middleware pre ")
		c.Next()
		fmt.Println("middleware post ")
		return nil
	}
}
