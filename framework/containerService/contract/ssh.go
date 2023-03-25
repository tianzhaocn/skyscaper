package contract

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"github.com/tianzhaocn/skyscraper/framework"
)

const SSHKey = "skyscraper:ssh"

// SSHOption 代表初始化的时候的选项
type SSHOption func(container framework.Container, config *SSHConfig) error

// SSHService 表示一个ssh服务
type SSHService interface {
	// GetClient 获取ssh连接实例
	GetClient(option ...SSHOption) (*ssh.Client, error)
}

// SSHConfig 为skyscraper定义的SSH配置结构
type SSHConfig struct {
	NetWork string
	Host    string
	Port    string
	*ssh.ClientConfig
}

// UniqKey 用来唯一标识一个SSHConfig配置
func (config *SSHConfig) UniqKey() string {
	return fmt.Sprintf("%v_%v_%v", config.Host, config.Port, config.User)
}
