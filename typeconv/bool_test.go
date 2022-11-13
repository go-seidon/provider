package typeconv_test

import (
	"github.com/go-seidon/provider/typeconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Bool Package", func() {

	Context("Bool function", Label("unit"), func() {

		When("input is false", func() {
			It("should return false", func() {
				res := typeconv.Bool(false)

				expectRes := false
				Expect(res).To(Equal(&expectRes))
			})
		})

		When("input is true", func() {
			It("should return true", func() {
				res := typeconv.Bool(true)

				expectRes := true
				Expect(res).To(Equal(&expectRes))
			})
		})
	})

	Context("BoolVal function", Label("unit"), func() {
		When("input is nil", func() {
			It("should return false", func() {
				res := typeconv.BoolVal(nil)

				Expect(res).To(Equal(false))
			})
		})

		When("input is false", func() {
			It("should return false", func() {
				input := false
				res := typeconv.BoolVal(&input)

				Expect(res).To(Equal(false))
			})
		})

		When("input is true", func() {
			It("should return true", func() {
				input := true
				res := typeconv.BoolVal(&input)

				Expect(res).To(Equal(true))
			})
		})
	})

})
