package demoservice

import "github.com/tianzhaocn/skyscraper/framework"

// DemoProvider 提供App的具体实现方法
type DemoProvider struct {
	BaseFolder string
}

// Register 注册App方法
func (app *DemoProvider) Register(container framework.Container) framework.NewInstance {
	return NewDemo
}

// Boot 启动调用
func (app *DemoProvider) Boot(container framework.Container) error {
	return nil
}

// IsDefer 是否延迟初始化
func (app *DemoProvider) IsDefer() bool {
	return false
}

// Params 获取初始化参数
func (app *DemoProvider) Params(container framework.Container) []interface{} {
	return nil
}

// Name 获取字符串凭证
func (app *DemoProvider) Name() string {
	return DemoKey
}
