package go_zero_config

import (
	"github.com/go-packagist/framework/container"
	"github.com/go-packagist/framework/contracts/provider"
	"github.com/zeromicro/go-zero/core/conf"
)

type Provider struct {
	configFile string
	*container.Container
}

var _ provider.Provider = (*Provider)(nil)

func NewProvider(c *container.Container, configFile string) *Provider {
	return &Provider{
		configFile: configFile,
		Container:  c,
	}
}

func (p *Provider) Register() {
	var cfg *Config
	conf.MustLoad(p.configFile, &cfg)

	p.Instance("zero.config", cfg)
}

func (p *Provider) Boot() {
}
