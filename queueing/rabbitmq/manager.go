package rabbitmq

import (
	"context"
)

// @note: call init on app startup
func (que *rabbitQueue) Open(ctx context.Context) error {
	return que.conn.Init(ctx)
}

func (que *rabbitQueue) Close(ctx context.Context) error {
	return que.conn.Close()
}
