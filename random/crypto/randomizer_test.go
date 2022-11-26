package crypto_test

import (
	"fmt"
	"testing"

	"github.com/go-seidon/provider/random"
	"github.com/go-seidon/provider/random/crypto"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCrypto(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Crypto Package")
}

var _ = Describe("Randomizer Package", func() {
	Context("String function", Label("unit"), func() {
		var (
			r random.Randomizer
		)

		BeforeEach(func() {
			r = crypto.NewRandomizer()
		})

		When("function is called one time", func() {
			It("should return result", func() {
				res, err := r.String(5)

				Expect(len(res)).To(Equal(5))
				Expect(err).To(BeNil())
			})
		})

		When("function is called multiple time", func() {
			It("should return result", func() {
				res1, err1 := r.String(5)
				res2, err2 := r.String(5)
				res3, err3 := r.String(5)

				Expect(len(res1)).To(Equal(5))
				Expect(err1).To(BeNil())

				Expect(len(res2)).To(Equal(5))
				Expect(err2).To(BeNil())

				Expect(len(res3)).To(Equal(5))
				Expect(err3).To(BeNil())

				Expect(res1).ToNot(Equal(res2))
				Expect(res2).ToNot(Equal(res3))
				Expect(res3).ToNot(Equal(res1))
			})
		})

		When("dictionary is empty", func() {
			It("should return error", func() {
				r := crypto.NewRandomizer(
					crypto.WithDictionary(""),
				)
				res, err := r.String(5)

				Expect(res).To(Equal(""))
				Expect(err).To(Equal(fmt.Errorf("failed generate random string")))
			})
		})
	})
})
