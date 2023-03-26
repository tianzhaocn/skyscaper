package main

import (
	"github.com/tianzhaocn/skyscraper/app/console"
	"github.com/tianzhaocn/skyscraper/app/http"
	"github.com/tianzhaocn/skyscraper/app/provider"
	"github.com/tianzhaocn/skyscraper/framework"
	"github.com/tianzhaocn/skyscraper/framework/containerService/service/app"
	"github.com/tianzhaocn/skyscraper/framework/containerService/service/config"
	"github.com/tianzhaocn/skyscraper/framework/containerService/service/env"
	"github.com/tianzhaocn/skyscraper/framework/containerService/service/id"
	"github.com/tianzhaocn/skyscraper/framework/containerService/service/kernel"
	"github.com/tianzhaocn/skyscraper/framework/containerService/service/log"
	"github.com/tianzhaocn/skyscraper/framework/containerService/service/orm"
	"github.com/tianzhaocn/skyscraper/framework/containerService/service/ssh"
)

func main() {
	// 初始化服务容器
	container := framework.NewSkyscraperContainer()
	// 绑定App服务提供者
	container.Bind(&app.AppProvider{})
	// 后续初始化需要绑定的服务提供者...
	container.Bind(&env.EnvProvider{})
	container.Bind(&config.ConfigProvider{})
	container.Bind(&id.IDProvider{})
	container.Bind(&log.LogServiceProvider{})
	container.Bind(&orm.GormProvider{})
	container.Bind(&ssh.SSHProvider{})
	// 将HTTP引擎初始化,并且作为服务提供者绑定到服务容器中
	
	provider.BindMyService(container)
	if engine, err := http.NewHttpEngine(container); err == nil {
		container.Bind(&kernel.KernelProvider{HttpEngine: engine})
	}
	console.RunCommand(container)
}
