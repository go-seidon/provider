package logrus_test

import (
	"github.com/go-seidon/provider/logging/logrus"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Optiona Package", func() {

	Context("WithAppContext function", Label("unit"), func() {
		When("parameter is specified", func() {
			It("should return result", func() {
				opt := logrus.WithAppContext("mock-name", "mock-version")
				var res logrus.LogParam
				opt(&res)

				Expect(res.AppName).To(Equal("mock-name"))
				Expect(res.AppVersion).To(Equal("mock-version"))
				Expect(res.DebuggingEnabled).To(BeFalse())
				Expect(res.PrettyPrintEnabled).To(BeFalse())
				Expect(res.StackSkip).To(BeNil())
			})
		})
	})

	Context("EnableDebugging function", Label("unit"), func() {
		When("function is called", func() {
			It("should return result", func() {
				opt := logrus.EnableDebugging()
				var res logrus.LogParam
				opt(&res)

				Expect(res.DebuggingEnabled).To(BeTrue())
				Expect(res.AppName).To(Equal(""))
				Expect(res.AppVersion).To(Equal(""))
				Expect(res.PrettyPrintEnabled).To(BeFalse())
				Expect(res.StackSkip).To(BeNil())
			})
		})
	})

	Context("EnablePrettyPrint function", Label("unit"), func() {
		When("function is called", func() {
			It("should return result", func() {
				opt := logrus.EnablePrettyPrint()
				var res logrus.LogParam
				opt(&res)

				Expect(res.PrettyPrintEnabled).To(BeTrue())
				Expect(res.DebuggingEnabled).To(BeFalse())
				Expect(res.AppName).To(Equal(""))
				Expect(res.AppVersion).To(Equal(""))
				Expect(res.StackSkip).To(BeNil())
			})
		})
	})

	Context("AddStackSkip function", Label("unit"), func() {
		When("add one stack skip", func() {
			It("should return result", func() {
				opt := logrus.AddStackSkip("some-pkg")
				var res logrus.LogParam
				opt(&res)

				Expect(res.StackSkip).To(Equal([]string{
					"some-pkg",
				}))
				Expect(res.PrettyPrintEnabled).To(BeFalse())
				Expect(res.DebuggingEnabled).To(BeFalse())
				Expect(res.AppName).To(Equal(""))
				Expect(res.AppVersion).To(Equal(""))
			})
		})

		When("add two stack skip", func() {
			It("should return result", func() {
				opt1 := logrus.AddStackSkip("some-pkg-1")
				opt2 := logrus.AddStackSkip("some-pkg-2")
				var res logrus.LogParam
				opt1(&res)
				opt2(&res)

				Expect(res.StackSkip).To(Equal([]string{
					"some-pkg-1",
					"some-pkg-2",
				}))
				Expect(res.PrettyPrintEnabled).To(BeFalse())
				Expect(res.DebuggingEnabled).To(BeFalse())
				Expect(res.AppName).To(Equal(""))
				Expect(res.AppVersion).To(Equal(""))
			})
		})
	})

})
