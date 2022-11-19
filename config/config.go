package config

import (
	"time"
)

type Config interface {
	Getter
	Setter
	Checker
	Manager
}

type Getter interface {
	Get(key string) (interface{}, error)
	GetBool(key string) (bool, error)
	GetFloat64(key string) (float64, error)
	GetInt(key string) (int, error)
	GetString(key string) (string, error)
	GetTime(key string) (time.Time, error)
	GetDuration(key string) (time.Duration, error)
}

type Setter interface {
	Set(key string, value interface{}) error
	SetDefault(key string, value interface{}) error
}

type Checker interface {
	IsSet(key string) (bool, error)
}

// @note: deprecate soon since this func is probably doing too much
type Manager interface {
	LoadConfig() error
	ParseConfig(cfg interface{}) error
}
