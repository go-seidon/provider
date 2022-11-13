package typeconv_test

import (
	"github.com/go-seidon/provider/typeconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Float Package", func() {

	Context("Float32 function", Label("unit"), func() {
		When("input is positive", func() {
			It("should return positive", func() {
				res := typeconv.Float32(2)

				expectRes := float32(2)
				Expect(res).To(Equal(&expectRes))
			})
		})

		When("input is negative", func() {
			It("should return negative", func() {
				res := typeconv.Float32(-2)

				expectRes := float32(-2)
				Expect(res).To(Equal(&expectRes))
			})
		})

		When("input is zero", func() {
			It("should return zero", func() {
				res := typeconv.Float32(0)

				expectRes := float32(0)
				Expect(res).To(Equal(&expectRes))
			})
		})
	})

	Context("Float32Val function", Label("unit"), func() {
		When("input is nil", func() {
			It("should return zero", func() {
				res := typeconv.Float32Val(nil)

				Expect(res).To(Equal(float32(0)))
			})
		})

		When("input is positive", func() {
			It("should return positive", func() {
				input := float32(2)
				res := typeconv.Float32Val(&input)

				Expect(res).To(Equal(float32(2)))
			})
		})

		When("input is negative", func() {
			It("should return negative", func() {
				input := float32(-2)
				res := typeconv.Float32Val(&input)

				Expect(res).To(Equal(float32(-2)))
			})
		})

		When("input is zero", func() {
			It("should return zero", func() {
				input := float32(0)
				res := typeconv.Float32Val(&input)

				Expect(res).To(Equal(float32(0)))
			})
		})
	})

	Context("Float64 function", Label("unit"), func() {
		When("input is positive", func() {
			It("should return positive", func() {
				res := typeconv.Float64(2)

				expectRes := float64(2)
				Expect(res).To(Equal(&expectRes))
			})
		})

		When("input is negative", func() {
			It("should return negative", func() {
				res := typeconv.Float64(-2)

				expectRes := float64(-2)
				Expect(res).To(Equal(&expectRes))
			})
		})

		When("input is zero", func() {
			It("should return zero", func() {
				res := typeconv.Float64(0)

				expectRes := float64(0)
				Expect(res).To(Equal(&expectRes))
			})
		})
	})

	Context("Float64Val function", Label("unit"), func() {
		When("input is nil", func() {
			It("should return zero", func() {
				res := typeconv.Float64Val(nil)

				Expect(res).To(Equal(float64(0)))
			})
		})

		When("input is positive", func() {
			It("should return positive", func() {
				input := float64(2)
				res := typeconv.Float64Val(&input)

				Expect(res).To(Equal(float64(2)))
			})
		})

		When("input is negative", func() {
			It("should return negative", func() {
				input := float64(-2)
				res := typeconv.Float64Val(&input)

				Expect(res).To(Equal(float64(-2)))
			})
		})

		When("input is zero", func() {
			It("should return zero", func() {
				input := float64(0)
				res := typeconv.Float64Val(&input)

				Expect(res).To(Equal(float64(0)))
			})
		})
	})

})
