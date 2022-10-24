package identifier_test

import (
	"testing"

	"github.com/go-seidon/provider/identifier"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestText(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Text Package")
}

var _ = Describe("KSU Identifier", func() {

	Context("GenerateId function", Label("unit"), func() {
		When("success generate id", func() {
			It("should return result", func() {
				ksuid := identifier.NewKsuid()
				id, err := ksuid.GenerateId()

				Expect(err).To(BeNil())
				Expect(id).ToNot(BeEmpty())
			})
		})
	})

})
