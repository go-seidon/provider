package mysql

type ClientOption = func(*ClientParam)

type ClientParam struct {
	Host            string
	Port            int
	Username        string
	Password        string
	DbName          string
	ShouldParseTime bool
}

type ClientConfig struct {
	DbName string
}

func ParseTime() ClientOption {
	return func(p *ClientParam) {
		p.ShouldParseTime = true
	}
}

func WithLocation(host string, port int) ClientOption {
	return func(p *ClientParam) {
		p.Host = host
		p.Port = port
	}
}

func WithAuth(username, password string) ClientOption {
	return func(p *ClientParam) {
		p.Username = username
		p.Password = password
	}
}

func WithConfig(cfg ClientConfig) ClientOption {
	return func(p *ClientParam) {
		p.DbName = cfg.DbName
	}
}
