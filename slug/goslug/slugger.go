package goslug

import (
	"github.com/gosimple/slug"
)

type goSlugger struct {
}

func (s *goSlugger) GenerateSlug(t string) string {
	return slug.Make(t)
}

func NewSlugger() *goSlugger {
	return &goSlugger{}
}
