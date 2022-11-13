package typeconv_test

import (
	"github.com/go-seidon/provider/typeconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("String Package", func() {

	Context("String function", Label("unit"), func() {

		When("input is empty", func() {
			It("should return empty", func() {
				res := typeconv.String("")

				expectRes := ""
				Expect(res).To(Equal(&expectRes))
			})
		})

		When("input is not empty", func() {
			It("should return not empty", func() {
				res := typeconv.String("not empty")

				expectRes := "not empty"
				Expect(res).To(Equal(&expectRes))
			})
		})
	})

	Context("StringVal function", Label("unit"), func() {
		When("input is nil", func() {
			It("should return empty", func() {
				res := typeconv.StringVal(nil)

				Expect(res).To(Equal(""))
			})
		})

		When("input is empty", func() {
			It("should return empty", func() {
				input := ""
				res := typeconv.StringVal(&input)

				Expect(res).To(Equal(""))
			})
		})

		When("input is not empty", func() {
			It("should return not empty", func() {
				input := "not empty"
				res := typeconv.StringVal(&input)

				Expect(res).To(Equal("not empty"))
			})
		})
	})

})
