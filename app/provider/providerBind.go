package provider

import (
	"github.com/tianzhaocn/skyscraper/app/provider/demoservice"
	"github.com/tianzhaocn/skyscraper/framework"
)

func BindMyService(container framework.Container) {
	container.Bind(&demoservice.DemoProvider{})
}
