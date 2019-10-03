package config

import (
	"github.com/spf13/viper"
)

type Conf struct {
	config Config
	isRead bool
}

type Config struct {
	DB
	GRPC
	AMPQ
}

type DB struct {
	Dns string
}

type GRPC struct {
	Addr string
}

type AMPQ struct {
	Addr        string
	NotifyQueue string
}

func (c *Conf) GetConfig() (Config, error) {

	if !c.isRead {
		err := c.read()

		if err != nil {
			return c.config, err
		}

		c.isRead = true
	}

	return c.config, nil
}

func (c *Conf) read() error {
	viper.AddConfigPath("../../../../config")
	viper.AddConfigPath("config")

	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	return c.unmarshal()
}

func (c *Conf) unmarshal() error {
	err := viper.Unmarshal(&c.config)
	return err
}
