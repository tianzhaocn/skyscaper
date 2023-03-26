package http

import (
	"github.com/tianzhaocn/skyscraper/app/http/module/demo"
	"github.com/tianzhaocn/skyscraper/framework"
)

// Routes 绑定业务层路由
func Routes(r *framework.Core) {
	demo.Register(r)
}
