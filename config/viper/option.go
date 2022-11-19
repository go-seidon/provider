package viper

type ConfigOption struct {
	FileName string
}

type Option func(*ConfigOption)

func WithFileName(name string) Option {
	return func(co *ConfigOption) {
		co.FileName = name
	}
}
