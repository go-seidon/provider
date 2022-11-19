package viper

type ConfigParam struct {
	FileName string
}

type ConfigOption func(*ConfigParam)

func WithFileName(name string) ConfigOption {
	return func(co *ConfigParam) {
		co.FileName = name
	}
}
