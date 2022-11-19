package logrus_test

import (
	"bytes"
	"fmt"

	"github.com/go-seidon/provider/logging/logrus"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	logrus_client "github.com/sirupsen/logrus"
)

var _ = Describe("Logrus Formater Package", func() {

	Context("Format function", Label("unit"), func() {
		var (
			formatter *logrus.GoFormatter
			entry     *logrus_client.Entry
		)

		BeforeEach(func() {
			formatter = &logrus.GoFormatter{}
			entry = &logrus_client.Entry{}
		})

		When("success format message", func() {
			It("should return result", func() {
				res, err := formatter.Format(entry)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})

		When("data contain error", func() {
			It("should return result", func() {
				entry.Data = logrus_client.Fields{
					"error": fmt.Errorf("some error"),
				}
				res, err := formatter.Format(entry)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})

		When("data contain field", func() {
			It("should return result", func() {
				entry.Data = logrus_client.Fields{
					"key": "value",
				}
				res, err := formatter.Format(entry)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})

		When("buffer is specified", func() {
			It("should return result", func() {
				entry.Buffer = &bytes.Buffer{}
				res, err := formatter.Format(entry)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})

		When("pretty print is used", func() {
			It("should return result", func() {
				formatter.PrettyPrint = true
				res, err := formatter.Format(entry)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

})
