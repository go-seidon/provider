package protobuf_test

import (
	"fmt"
	"testing"

	"github.com/go-seidon/provider/serialization"
	"github.com/go-seidon/provider/serialization/protobuf"
	"github.com/go-seidon/provider/testdata"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/protobuf/proto"
)

func TestProtobuf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Protobuf Package")
}

var _ = Describe("Serializer Package", func() {
	Context("NewSerializer function", Label("unit"), func() {
		When("function is called", func() {
			It("should return result", func() {
				res := protobuf.NewSerializer()

				Expect(res).ToNot(BeNil())
			})
		})
	})

	Context("Marshal function", Label("unit"), func() {
		var (
			serializer serialization.Serializer
		)

		BeforeEach(func() {
			serializer = protobuf.NewSerializer()
		})

		When("success encode data", func() {
			It("should return result", func() {
				d := &testdata.SimpleMessage{
					String_: "text",
				}
				res, err := serializer.Marshal(d)

				expected, _ := proto.Marshal(d)
				Expect(err).To(BeNil())
				Expect(res).To(Equal(expected))
			})
		})

		When("failed encode data", func() {
			It("should return error", func() {
				d := make(chan int)
				res, err := serializer.Marshal(d)

				Expect(err).To(Equal(fmt.Errorf("invalid message")))
				Expect(res).To(BeNil())
			})
		})

		When("data is nil", func() {
			It("should return result", func() {
				res, err := serializer.Marshal(nil)

				Expect(err).To(Equal(fmt.Errorf("invalid message")))
				Expect(res).To(BeNil())
			})
		})
	})

	Context("Unmarshal function", Label("unit"), func() {

		var (
			d          []byte
			serializer serialization.Serializer
		)

		BeforeEach(func() {
			serializer = protobuf.NewSerializer()
			b := &testdata.SimpleMessage{
				String_: "text",
			}
			d, _ = proto.Marshal(b)
		})

		When("success decode data", func() {
			It("should return result", func() {
				var res testdata.SimpleMessage
				err := serializer.Unmarshal(d, &res)

				Expect(err).To(BeNil())
				Expect(res.String_).To(Equal("text"))
			})
		})

		When("failed decode data", func() {
			It("should return error", func() {
				var res string
				d = []byte{}
				err := serializer.Unmarshal(d, &res)

				Expect(err).ToNot(BeNil())
			})
		})

		When("data is nil", func() {
			It("should return empty result", func() {
				var res testdata.SimpleMessage
				err := serializer.Unmarshal(nil, &res)

				Expect(err).To(BeNil())
				Expect(res.String_).To(Equal(""))
			})
		})
	})
})
