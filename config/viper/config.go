package viper

import (
	"time"

	"github.com/go-seidon/provider/config"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

type viperConfig struct {
	client *viper.Viper
}

func (c *viperConfig) Get(key string) (interface{}, error) {
	if !c.client.IsSet(key) {
		return nil, config.ErrNotFound
	}

	return c.client.Get(key), nil
}

func (c *viperConfig) GetBool(key string) (bool, error) {
	if !c.client.IsSet(key) {
		return false, config.ErrNotFound
	}

	return c.client.GetBool(key), nil
}

func (c *viperConfig) GetFloat64(key string) (float64, error) {
	if !c.client.IsSet(key) {
		return 0, config.ErrNotFound
	}

	return c.client.GetFloat64(key), nil
}

func (c *viperConfig) GetInt(key string) (int, error) {
	if !c.client.IsSet(key) {
		return 0, config.ErrNotFound
	}

	return c.client.GetInt(key), nil
}

func (c *viperConfig) GetString(key string) (string, error) {
	if !c.client.IsSet(key) {
		return "", config.ErrNotFound
	}

	return c.client.GetString(key), nil
}

func (c *viperConfig) GetTime(key string) (time.Time, error) {
	if !c.client.IsSet(key) {
		return time.Time{}, config.ErrNotFound
	}

	return c.client.GetTime(key), nil
}

func (c *viperConfig) GetDuration(key string) (time.Duration, error) {
	if !c.client.IsSet(key) {
		return 0, config.ErrNotFound
	}

	return c.client.GetDuration(key), nil
}

func (c *viperConfig) Set(key string, value interface{}) error {
	c.client.Set(key, value)
	return nil
}

func (c *viperConfig) SetDefault(key string, value interface{}) error {
	c.client.SetDefault(key, value)
	return nil
}

func (c *viperConfig) IsSet(key string) (bool, error) {
	return c.client.IsSet(key), nil
}

func (c *viperConfig) LoadConfig() error {
	return c.client.ReadInConfig()
}

// @note: deprecate soon since this func is probably doing too much
func (c *viperConfig) ParseConfig(cfg interface{}) error {
	return c.client.Unmarshal(cfg, func(dc *mapstructure.DecoderConfig) {
		dc.TagName = "env"
	})
}

func NewConfig(opts ...Option) (*viperConfig, error) {
	option := ConfigOption{}
	for _, opt := range opts {
		opt(&option)
	}

	client := viper.New()
	if option.FileName != "" {
		client.SetConfigFile(option.FileName)
	}
	client.AllowEmptyEnv(true)
	client.AutomaticEnv()

	c := &viperConfig{
		client: client,
	}
	return c, nil
}
