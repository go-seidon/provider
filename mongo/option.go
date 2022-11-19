package mongo

type ClientOption = func(*ClientParam)

type ClientParam struct {
	DbMode string
	DbName string

	AuthMode     string
	AuthSource   string
	AuthUser     string
	AuthPassword string

	StdHost string
	StdPort int

	RsName  string
	RsHosts []string
}

type ClientLocation struct {
	Host string
	Port int
}

type ClientConfig struct {
	DbName   string
	DbMode   string
	AuthMode string
}

func (p *ClientParam) ModeStandalone() bool {
	return p.DbMode == MODE_STANDALONE
}

func (p *ClientParam) ModeReplication() bool {
	return p.DbMode == MODE_REPLICATION
}

func (p *ClientParam) ModeSupported() bool {
	return p.ModeStandalone() || p.ModeReplication()
}

func (p *ClientParam) AuthBasic() bool {
	return p.AuthMode == AUTH_BASIC
}

func (p *ClientParam) AuthSupported() bool {
	return p.AuthBasic()
}

func UsingStandalone(host string, port int) ClientOption {
	return func(p *ClientParam) {
		p.DbMode = MODE_STANDALONE
		p.StdHost = host
		p.StdPort = port
	}
}

func UsingReplication(rsName string, rsHosts []string) ClientOption {
	return func(p *ClientParam) {
		p.DbMode = MODE_REPLICATION
		p.RsName = rsName
		p.RsHosts = rsHosts
	}
}

func WithBasicAuth(username, password, source string) ClientOption {
	return func(p *ClientParam) {
		p.AuthUser = username
		p.AuthPassword = password
		p.AuthSource = source
		p.AuthMode = AUTH_BASIC
	}
}

func WithConfig(cfg ClientConfig) ClientOption {
	return func(p *ClientParam) {
		if cfg.DbName != "" {
			p.DbName = cfg.DbName
		}
		if cfg.AuthMode != "" {
			p.AuthMode = cfg.AuthMode
		}
		if cfg.DbMode != "" {
			p.DbMode = cfg.DbMode
		}
	}
}
