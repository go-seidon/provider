package typeconv_test

import (
	"time"

	"github.com/go-seidon/provider/typeconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Duration Package", func() {
	Context("Duration function", Label("unit"), func() {
		When("input is positive", func() {
			It("should return positive", func() {
				res := typeconv.Duration(time.Duration(2))

				expectRes := time.Duration(2)
				Expect(res).To(Equal(&expectRes))
			})
		})

		When("input is negative", func() {
			It("should return negative", func() {
				res := typeconv.Duration(time.Duration(-2))

				expectRes := time.Duration(-2)
				Expect(res).To(Equal(&expectRes))
			})
		})

		When("input is zero", func() {
			It("should return zero", func() {
				res := typeconv.Duration(time.Duration(0))

				expectRes := time.Duration(0)
				Expect(res).To(Equal(&expectRes))
			})
		})
	})

	Context("DurationVal function", Label("unit"), func() {
		When("input is nil", func() {
			It("should return zero", func() {
				res := typeconv.DurationVal(nil)

				Expect(res).To(Equal(time.Duration(0)))
			})
		})

		When("input is positive", func() {
			It("should return positive", func() {
				input := time.Duration(2)
				res := typeconv.DurationVal(&input)

				Expect(res).To(Equal(time.Duration(2)))
			})
		})

		When("input is negative", func() {
			It("should return negative", func() {
				input := time.Duration(-2)
				res := typeconv.DurationVal(&input)

				Expect(res).To(Equal(time.Duration(-2)))
			})
		})

		When("input is zero", func() {
			It("should return zero", func() {
				input := time.Duration(0)
				res := typeconv.DurationVal(&input)

				Expect(res).To(Equal(time.Duration(0)))
			})
		})
	})
})
