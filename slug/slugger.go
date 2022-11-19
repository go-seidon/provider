package slug

type Slugger interface {
	GenerateSlug(t string) string
}
