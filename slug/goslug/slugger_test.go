package goslug_test

import (
	"testing"

	"github.com/go-seidon/provider/slug/goslug"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGoslugger(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goslugger Package")
}

var _ = Describe("Slugger Package", func() {
	Context("NewSlugger function", Label("unit"), func() {
		When("function is called", func() {
			It("should return result", func() {
				res := goslug.NewSlugger()

				Expect(res).ToNot(BeNil())
			})
		})
	})

	Context("GenerateSlug function", Label("unit"), func() {
		When("function is called", func() {
			It("should return result", func() {
				slugger := goslug.NewSlugger()
				res := slugger.GenerateSlug("Random text")

				Expect(res).To(Equal("random-text"))
			})
		})
	})
})
