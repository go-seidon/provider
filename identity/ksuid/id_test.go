package ksuid_test

import (
	"testing"

	"github.com/go-seidon/provider/identity/ksuid"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestKsuid(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ksuid Package")
}

var _ = Describe("KSU Identifier", func() {

	Context("GenerateId function", Label("unit"), func() {
		When("success generate id", func() {
			It("should return result", func() {
				ksuid := ksuid.NewIdentifier()
				id, err := ksuid.GenerateId()

				Expect(err).To(BeNil())
				Expect(id).ToNot(BeEmpty())
			})
		})
	})

})
