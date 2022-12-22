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

				r := time.Time{}
				Expect(res).To(Equal(&r))
			})
		})

		When("input is non empty", func() {
			It("should return non empty", func() {
				currentTs := time.Now().UTC()
				res := typeconv.Time(currentTs)

				r := currentTs
				Expect(res).To(Equal(&r))
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
				input := time.Time{}.UTC()
				res := typeconv.TimeVal(&input)

				Expect(res).To(Equal(time.Time{}))
			})
		})

		When("input is non empty", func() {
			It("should return non empty", func() {
				input := time.Now().UTC()
				res := typeconv.TimeVal(&input)

				Expect(res).To(Equal(input))
			})
		})
	})

	Context("UnixMilli function", Label("unit"), func() {
		When("input is nil", func() {
			It("should return nil", func() {
				res := typeconv.UnixMilli(nil)

				Expect(res).To(BeNil())
			})
		})

		When("input is empty", func() {
			It("should return empty", func() {
				input := int64(0)
				res := typeconv.UnixMilli(&input)

				r := time.UnixMilli(input).UTC()
				Expect(res).To(Equal(&r))
			})
		})

		When("input is non empty", func() {
			It("should return non empty", func() {
				ts := time.Now()
				input := ts.UnixMilli()
				res := typeconv.UnixMilli(&input)

				r := time.UnixMilli(ts.UnixMilli()).UTC()
				Expect(res).To(Equal(&r))
			})
		})
	})

	Context("TimeMilli function", Label("unit"), func() {
		When("input is nil", func() {
			It("should return nil", func() {
				res := typeconv.TimeMilli(nil)

				Expect(res).To(BeNil())
			})
		})

		When("input is empty", func() {
			It("should return empty", func() {
				input := time.UnixMilli(0)
				res := typeconv.TimeMilli(&input)

				r := input.UnixMilli()
				Expect(res).To(Equal(&r))
			})
		})

		When("input is non empty", func() {
			It("should return non empty", func() {
				input := time.Now()
				res := typeconv.TimeMilli(&input)

				r := input.UnixMilli()
				Expect(res).To(Equal(&r))
			})
		})
	})
})
