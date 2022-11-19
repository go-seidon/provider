package http

type ClientOption = func(*ClientParam)

type ClientParam struct {
	ShouldCheckSSL bool
}

func ShouldCheckSSL() ClientOption {
	return func(p *ClientParam) {
		p.ShouldCheckSSL = true
	}
}
