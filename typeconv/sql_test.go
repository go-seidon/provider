package typeconv_test

import (
	"database/sql"
	"time"

	"github.com/go-seidon/provider/typeconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sql Package", func() {
	Context("SqlBool function", Label("unit"), func() {
		When("input is invalid", func() {
			It("should return nil", func() {
				input := sql.NullBool{}
				res := typeconv.SqlBool(input)

				Expect(res).To(BeNil())
			})
		})

		When("input is true", func() {
			It("should return result", func() {
				input := sql.NullBool{
					Bool:  true,
					Valid: true,
				}
				res := typeconv.SqlBool(input)

				r := true
				Expect(res).To(Equal(&r))
			})
		})

		When("input is false", func() {
			It("should return result", func() {
				input := sql.NullBool{
					Bool:  false,
					Valid: true,
				}
				res := typeconv.SqlBool(input)

				r := false
				Expect(res).To(Equal(&r))
			})
		})
	})

	Context("SqlBoolVal function", Label("unit"), func() {
		When("input is nil", func() {
			It("should return empty", func() {
				res := typeconv.SqlBoolVal(nil)

				Expect(res).To(Equal(sql.NullBool{}))
			})
		})

		When("input is false", func() {
			It("should return result", func() {
				input := false
				res := typeconv.SqlBoolVal(&input)

				Expect(res).To(Equal(sql.NullBool{
					Bool:  false,
					Valid: true,
				}))
			})
		})

		When("input is true", func() {
			It("should return true", func() {
				input := true
				res := typeconv.SqlBoolVal(&input)

				Expect(res).To(Equal(sql.NullBool{
					Bool:  true,
					Valid: true,
				}))
			})
		})
	})

	Context("SqlFloat64 function", Label("unit"), func() {
		When("input is invalid", func() {
			It("should return nil", func() {
				input := sql.NullFloat64{}
				res := typeconv.SqlFloat64(input)

				Expect(res).To(BeNil())
			})
		})

		When("input is valid", func() {
			It("should return result", func() {
				input := sql.NullFloat64{
					Float64: 32.335,
					Valid:   true,
				}
				res := typeconv.SqlFloat64(input)

				r := 32.335
				Expect(res).To(Equal(&r))
			})
		})
	})

	Context("SqlFloat64Val function", Label("unit"), func() {
		When("input is nil", func() {
			It("should return empty", func() {
				res := typeconv.SqlFloat64Val(nil)

				Expect(res).To(Equal(sql.NullFloat64{}))
			})
		})

		When("input is zero", func() {
			It("should return result", func() {
				input := float64(0)
				res := typeconv.SqlFloat64Val(&input)

				Expect(res).To(Equal(sql.NullFloat64{
					Float64: 0,
					Valid:   true,
				}))
			})
		})

		When("input is non zero", func() {
			It("should return result", func() {
				input := 32.335
				res := typeconv.SqlFloat64Val(&input)

				Expect(res).To(Equal(sql.NullFloat64{
					Float64: 32.335,
					Valid:   true,
				}))
			})
		})
	})

	Context("SqlInt32 function", Label("unit"), func() {
		When("input is invalid", func() {
			It("should return nil", func() {
				input := sql.NullInt32{}
				res := typeconv.SqlInt32(input)

				Expect(res).To(BeNil())
			})
		})

		When("input is valid", func() {
			It("should return result", func() {
				input := sql.NullInt32{
					Int32: 32,
					Valid: true,
				}
				res := typeconv.SqlInt32(input)

				r := int32(32)
				Expect(res).To(Equal(&r))
			})
		})
	})

	Context("SqlInt32Val function", Label("unit"), func() {
		When("input is nil", func() {
			It("should return empty", func() {
				res := typeconv.SqlInt32Val(nil)

				Expect(res).To(Equal(sql.NullInt32{}))
			})
		})

		When("input is zero", func() {
			It("should return result", func() {
				input := int32(0)
				res := typeconv.SqlInt32Val(&input)

				Expect(res).To(Equal(sql.NullInt32{
					Int32: 0,
					Valid: true,
				}))
			})
		})

		When("input is non zero", func() {
			It("should return result", func() {
				input := int32(32)
				res := typeconv.SqlInt32Val(&input)

				Expect(res).To(Equal(sql.NullInt32{
					Int32: 32,
					Valid: true,
				}))
			})
		})
	})

	Context("SqlInt64 function", Label("unit"), func() {
		When("input is invalid", func() {
			It("should return nil", func() {
				input := sql.NullInt64{}
				res := typeconv.SqlInt64(input)

				Expect(res).To(BeNil())
			})
		})

		When("input is valid", func() {
			It("should return result", func() {
				input := sql.NullInt64{
					Int64: 64,
					Valid: true,
				}
				res := typeconv.SqlInt64(input)

				r := int64(64)
				Expect(res).To(Equal(&r))
			})
		})
	})

	Context("SqlInt64Val function", Label("unit"), func() {
		When("input is nil", func() {
			It("should return empty", func() {
				res := typeconv.SqlInt64Val(nil)

				Expect(res).To(Equal(sql.NullInt64{}))
			})
		})

		When("input is zero", func() {
			It("should return result", func() {
				input := int64(0)
				res := typeconv.SqlInt64Val(&input)

				Expect(res).To(Equal(sql.NullInt64{
					Int64: 0,
					Valid: true,
				}))
			})
		})

		When("input is non zero", func() {
			It("should return result", func() {
				input := int64(64)
				res := typeconv.SqlInt64Val(&input)

				Expect(res).To(Equal(sql.NullInt64{
					Int64: 64,
					Valid: true,
				}))
			})
		})
	})

	Context("SqlString function", Label("unit"), func() {
		When("input is invalid", func() {
			It("should return nil", func() {
				input := sql.NullString{}
				res := typeconv.SqlString(input)

				Expect(res).To(BeNil())
			})
		})

		When("input is valid", func() {
			It("should return result", func() {
				input := sql.NullString{
					String: "hello",
					Valid:  true,
				}
				res := typeconv.SqlString(input)

				r := "hello"
				Expect(res).To(Equal(&r))
			})
		})
	})

	Context("SqlStringVal function", Label("unit"), func() {
		When("input is nil", func() {
			It("should return empty", func() {
				res := typeconv.SqlStringVal(nil)

				Expect(res).To(Equal(sql.NullString{}))
			})
		})

		When("input is empty", func() {
			It("should return empty", func() {
				input := ""
				res := typeconv.SqlStringVal(&input)

				Expect(res).To(Equal(sql.NullString{
					String: "",
					Valid:  true,
				}))
			})
		})

		When("input is non empty", func() {
			It("should return result", func() {
				input := "hello"
				res := typeconv.SqlStringVal(&input)

				Expect(res).To(Equal(sql.NullString{
					String: "hello",
					Valid:  true,
				}))
			})
		})
	})

	Context("SqlTime function", Label("unit"), func() {
		When("input is invalid", func() {
			It("should return nil", func() {
				input := sql.NullTime{}
				res := typeconv.SqlTime(input)

				Expect(res).To(BeNil())
			})
		})

		When("input is valid", func() {
			It("should return result", func() {
				ts := time.Now()
				input := sql.NullTime{
					Time:  ts,
					Valid: true,
				}
				res := typeconv.SqlTime(input)

				r := ts.UTC()
				Expect(res).To(Equal(&r))
			})
		})
	})

	Context("SqlTimeVal function", Label("unit"), func() {
		When("input is nil", func() {
			It("should return empty", func() {
				res := typeconv.SqlTimeVal(nil)

				Expect(res).To(Equal(sql.NullTime{}))
			})
		})

		When("input is empty", func() {
			It("should return empty", func() {
				input := time.Time{}
				res := typeconv.SqlTimeVal(&input)

				Expect(res).To(Equal(sql.NullTime{
					Time:  input,
					Valid: true,
				}))
			})
		})

		When("input is non empty", func() {
			It("should return result", func() {
				input := time.Now()
				res := typeconv.SqlTimeVal(&input)

				Expect(res).To(Equal(sql.NullTime{
					Time:  input.UTC(),
					Valid: true,
				}))
			})
		})
	})

	Context("SqlUnixMilli function", Label("unit"), func() {
		When("input is invalid", func() {
			It("should return nil", func() {
				input := sql.NullInt64{}
				res := typeconv.SqlUnixMilli(input)

				Expect(res).To(BeNil())
			})
		})

		When("input is valid", func() {
			It("should return result", func() {
				ts := time.Now()
				input := sql.NullInt64{
					Int64: ts.UnixMilli(),
					Valid: true,
				}
				res := typeconv.SqlUnixMilli(input)

				r := time.UnixMilli(ts.UnixMilli()).UTC()
				Expect(res).To(Equal(&r))
			})
		})
	})

	Context("SqlTimeMilli function", Label("unit"), func() {
		When("input is nil", func() {
			It("should return empty", func() {
				res := typeconv.SqlTimeMilli(nil)

				Expect(res).To(Equal(sql.NullInt64{}))
			})
		})

		When("input is valid", func() {
			It("should return result", func() {
				input := time.Now()
				res := typeconv.SqlTimeMilli(&input)

				Expect(res).To(Equal(sql.NullInt64{
					Int64: input.UnixMilli(),
					Valid: true,
				}))
			})
		})
	})
})
