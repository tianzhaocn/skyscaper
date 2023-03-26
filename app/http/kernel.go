package http

import (
	
	"github.com/tianzhaocn/skyscraper/framework"
	"github.com/tianzhaocn/skyscraper/framework/middleware"
)

// NewHttpCore 创建了一个绑定了路由的Web引擎
func NewHttpEngine(container framework.Container) (*framework.Core, error) {
	core := framework.NewCore()
	// 设置了Engine
	core.SetContainer(container)

	// 默认注册recovery中间件
	core.Use(middleware.Recovery())

	// 业务绑定路由操作
	Routes(core)
	// 返回绑定路由后的Web引擎
	return core, nil
}
