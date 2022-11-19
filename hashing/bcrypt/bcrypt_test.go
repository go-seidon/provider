package bcrypt_test

import (
	"fmt"
	"testing"

	"github.com/go-seidon/provider/hashing"
	"github.com/go-seidon/provider/hashing/bcrypt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBcrypt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bcrypt Package")
}

var _ = Describe("Bcrypt Hasher Package", func() {
	Context("NewHasher function", Label("unit"), func() {
		When("function is called", func() {
			It("should return result", func() {
				res := bcrypt.NewHasher()

				Expect(res).ToNot(BeNil())
			})
		})
	})

	Context("Generate function", Label("unit", "slow"), func() {
		var (
			h    hashing.Hasher
			text string
		)

		BeforeEach(func() {
			h = bcrypt.NewHasher()
			text = "some-secret"
		})

		When("success generate hash", func() {
			It("should return result", func() {
				res, err := h.Generate(text)

				equalIfNil := h.Verify(string(res), text)

				Expect(res).ToNot(BeEmpty())
				Expect(err).To(BeNil())
				Expect(equalIfNil).To(BeNil())
			})
		})
	})

	Context("Verify function", Label("unit", "slow"), func() {
		var (
			h    hashing.Hasher
			text string
			hash string
		)

		BeforeEach(func() {
			h = bcrypt.NewHasher()
			text = "some-secret"
			hash = "$2a$10$xA9.FPfIYi2ZI6V5/jw5leFVUCjsgN4lBS5iS8loLv1hngJj1ys/2"
		})

		When("hash is equal", func() {
			It("should return nil", func() {
				err := h.Verify(hash, text)

				Expect(err).To(BeNil())
			})
		})

		When("hash is not equal", func() {
			It("should return nil", func() {
				err := h.Verify(hash, "other-secret")

				Expect(err).To(Equal(fmt.Errorf("crypto/bcrypt: hashedPassword is not the hash of the given password")))
			})
		})
	})

})
