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
	return func(lo *LogParam) {
		lo.AppCtxEnabled = true
		lo.AppName = name
		lo.AppVersion = version
	}
}

func EnableDebugging() LogOption {
	return func(lo *LogParam) {
		lo.DebuggingEnabled = true
	}
}

func EnablePrettyPrint() LogOption {
	return func(lo *LogParam) {
		lo.PrettyPrintEnabled = true
	}
}

func AddStackSkip(pkg string) LogOption {
	return func(lo *LogParam) {
		lo.StackSkip = append(lo.StackSkip, pkg)
	}
}
