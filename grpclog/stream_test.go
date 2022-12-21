package grpclog_test

import (
	"fmt"

	"github.com/go-seidon/provider/grpc"
	mock_grpc "github.com/go-seidon/provider/grpc/mock"
	"github.com/go-seidon/provider/grpclog"
	mock_logging "github.com/go-seidon/provider/logging/mock"
	"github.com/go-seidon/provider/testdata"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stream Package", func() {

	Describe("Log Server Stream", func() {
		var (
			lss    grpc.ServerStream
			ss     *mock_grpc.MockServerStream
			logger *mock_logging.MockLogger
			msg    *testdata.SimpleMessage
		)

		BeforeEach(func() {
			t := GinkgoT()
			ctrl := gomock.NewController(t)
			ss = mock_grpc.NewMockServerStream(ctrl)
			logger = mock_logging.NewMockLogger(ctrl)
			lss = grpclog.NewLogServerStream(ss, logger)
			msg = &testdata.SimpleMessage{}
		})

		Context("SendMsg function", Label("unit"), func() {
			When("error happened", func() {
				It("should return error", func() {
					expectErr := fmt.Errorf("network error")

					ss.
						EXPECT().
						SendMsg(gomock.Eq(msg)).
						Return(expectErr).
						Times(1)

					err := lss.SendMsg(msg)

					Expect(err).To(Equal(expectErr))
				})
			})

			When("error not happened", func() {
				It("should return result", func() {
					ss.
						EXPECT().
						SendMsg(gomock.Eq(msg)).
						Return(nil).
						Times(1)

					logger.
						EXPECT().
						WithFields(gomock.Any()).
						Return(logger).
						Times(1)

					logger.
						EXPECT().
						Info(gomock.Eq("send stream")).
						Times(1)

					err := lss.SendMsg(msg)

					Expect(err).To(BeNil())
				})
			})
		})

		Context("RecvMsg function", Label("unit"), func() {
			When("error happened", func() {
				It("should return error", func() {
					expectErr := fmt.Errorf("network error")

					ss.
						EXPECT().
						RecvMsg(gomock.Eq(msg)).
						Return(expectErr).
						Times(1)

					err := lss.RecvMsg(msg)

					Expect(err).To(Equal(expectErr))
				})
			})

			When("error not happened", func() {
				It("should return result", func() {
					ss.
						EXPECT().
						RecvMsg(gomock.Eq(msg)).
						Return(nil).
						Times(1)

					logger.
						EXPECT().
						WithFields(gomock.Any()).
						Return(logger).
						Times(1)

					logger.
						EXPECT().
						Info(gomock.Eq("receive stream")).
						Times(1)

					err := lss.RecvMsg(msg)

					Expect(err).To(BeNil())
				})
			})
		})
	})

})
