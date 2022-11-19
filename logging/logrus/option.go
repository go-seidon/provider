package logrus

type LogParam struct {
	AppCtxEnabled bool
	AppName       string
	AppVersion    string

	DebuggingEnabled   bool
	PrettyPrintEnabled bool

	StackSkip []string
}

type LogOption func(*LogParam)

func WithAppContext(name, version string) LogOption {
	return func(p *LogParam) {
		p.AppCtxEnabled = true
		p.AppName = name
		p.AppVersion = version
	}
}

func EnableDebugging() LogOption {
	return func(p *LogParam) {
		p.DebuggingEnabled = true
	}
}

func EnablePrettyPrint() LogOption {
	return func(p *LogParam) {
		p.PrettyPrintEnabled = true
	}
}

func AddStackSkip(pkg string) LogOption {
	return func(p *LogParam) {
		p.StackSkip = append(p.StackSkip, pkg)
	}
}
