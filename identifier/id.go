package identifier

type Identifier interface {
	GenerateId() (string, error)
}
