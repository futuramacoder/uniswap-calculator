package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Environment struct {
	AppName    string `envconfig:"APP_NAME" default:"uniswap-calculator"`
	EthNodeUrl string `envconfig:"ETH_NODE_URL"`
}

func LoadEnvConfig() (*Environment, error) {
	e := &Environment{}
	if err := envconfig.Process("", e); err != nil {
		return nil, err
	}
	return e, nil
}
