package identity

type Identifier interface {
	GenerateId() (string, error)
}
