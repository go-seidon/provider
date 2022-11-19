package base64_test

import (
	encoding_base64 "encoding/base64"
	"testing"

	"github.com/go-seidon/provider/encoding"
	"github.com/go-seidon/provider/encoding/base64"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBase64(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Base64 Package")
}

var _ = Describe("Base64 Encoder Package", func() {
	Context("NewEncoder function", Label("unit"), func() {
		When("function is called", func() {
			It("should return result", func() {
				res := base64.NewEncoder()

				Expect(res).ToNot(BeNil())
			})
		})
	})

	Context("Encode function", Label("unit"), func() {
		var (
			encoder encoding.Encoder
			d       []byte
		)

		BeforeEach(func() {
			encoder = base64.NewEncoder()
			d = []byte("some-data")
		})

		When("success encode data", func() {
			It("should return result", func() {
				res, err := encoder.Encode(d)

				Expect(res).To(Equal("c29tZS1kYXRh"))
				Expect(err).To(BeNil())
			})
		})
	})

	Context("Decode function", Label("unit"), func() {
		var (
			encoder encoding.Encoder
			d       string
		)

		BeforeEach(func() {
			encoder = base64.NewEncoder()
			d = "c29tZS1kYXRh"
		})

		When("failed decode data", func() {
			It("should return error", func() {
				res, err := encoder.Decode("\\")

				Expect(res).To(Equal([]byte{}))
				Expect(err).To(Equal(encoding_base64.CorruptInputError(0)))
			})
		})

		When("success decode data", func() {
			It("should return result", func() {
				res, err := encoder.Decode(d)

				Expect(string(res)).To(Equal("some-data"))
				Expect(err).To(BeNil())
			})
		})
	})

})
