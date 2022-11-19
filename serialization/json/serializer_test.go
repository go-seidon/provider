package json_test

import (
	encoding_json "encoding/json"
	"testing"

	"github.com/go-seidon/provider/serialization"
	"github.com/go-seidon/provider/serialization/json"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestJson(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Json Package")
}

var _ = Describe("Serializer Package", func() {
	Context("NewSerializer function", Label("unit"), func() {
		When("function is called", func() {
			It("should return result", func() {
				res := json.NewSerializer()

				Expect(res).ToNot(BeNil())
			})
		})
	})

	Context("Marshal function", Label("unit"), func() {
		var (
			serializer serialization.Serializer
		)

		BeforeEach(func() {
			serializer = json.NewSerializer()
		})

		When("success encode data", func() {
			It("should return result", func() {
				d := struct{}{}
				res, err := serializer.Marshal(d)

				expected, _ := encoding_json.Marshal(d)
				Expect(err).To(BeNil())
				Expect(res).To(Equal(expected))
			})
		})

		When("failed encode data", func() {
			It("should return error", func() {
				d := make(chan int)
				res, err := serializer.Marshal(d)

				Expect(err).ToNot(BeNil())
				Expect(res).To(BeNil())
			})
		})

		When("data is nil", func() {
			It("should return result", func() {
				res, err := serializer.Marshal(nil)

				expected, _ := encoding_json.Marshal(nil)
				Expect(err).To(BeNil())
				Expect(res).To(Equal(expected))
			})
		})
	})

	Context("Unmarshal function", Label("unit"), func() {
		type Data struct {
			Val string `json:"val"`
		}

		var (
			d          []byte
			serializer serialization.Serializer
		)

		BeforeEach(func() {
			serializer = json.NewSerializer()
			d = []byte(`{"val":"ok"}`)
		})

		When("success decode data", func() {
			It("should return result", func() {
				var res Data
				err := serializer.Unmarshal(d, &res)

				Expect(err).To(BeNil())
				Expect(res.Val).To(Equal("ok"))
			})
		})

		When("failed decode data", func() {
			It("should return error", func() {
				var res Data
				d = []byte{}
				err := serializer.Unmarshal(d, &res)

				Expect(err).ToNot(BeNil())
				Expect(res.Val).To(Equal(""))
			})
		})

		When("data is nil", func() {
			It("should return result", func() {
				var res Data
				err := serializer.Unmarshal(nil, &res)

				Expect(err.Error()).To(Equal("unexpected end of JSON input"))
				Expect(res.Val).To(Equal(""))
			})
		})
	})
})
