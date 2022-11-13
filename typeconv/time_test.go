package typeconv_test

import (
	"time"

	"github.com/go-seidon/provider/typeconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Time Package", func() {

	Context("Time function", Label("unit"), func() {
		When("input is empty", func() {
			It("should return empty", func() {
				res := typeconv.Time(time.Time{})

				expectRes := time.Time{}
				Expect(res).To(Equal(&expectRes))
			})
		})

		When("input is non empty", func() {
			It("should return non empty", func() {
				currentTs := time.Now()
				res := typeconv.Time(currentTs)

				expectRes := currentTs
				Expect(res).To(Equal(&expectRes))
			})
		})
	})

	Context("TimeVal function", Label("unit"), func() {
		When("input is nil", func() {
			It("should return empty", func() {
				res := typeconv.TimeVal(nil)

				Expect(res).To(Equal(time.Time{}))
			})
		})

		When("input is empty", func() {
			It("should return empty", func() {
				input := time.Time{}
				res := typeconv.TimeVal(&input)

				Expect(res).To(Equal(time.Time{}))
			})
		})

		When("input is non empty", func() {
			It("should return non empty", func() {
				input := time.Now()
				res := typeconv.TimeVal(&input)

				Expect(res).To(Equal(input))
			})
		})
	})

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
