package crypto

type RandomizerParam struct {
	Dictionary string
}

type RandomizerOption = func(*RandomizerParam)

func WithDictionary(d string) RandomizerOption {
	return func(p *RandomizerParam) {
		p.Dictionary = d
	}
}
