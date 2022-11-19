package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	MODE_STANDALONE  = "standalone"
	MODE_REPLICATION = "replication"

	AUTH_BASIC = "basic"
)

type Client interface {
	Connect(ctx context.Context) error
	Ping(ctx context.Context, rp *readpref.ReadPref) error
	Database(name string, opts ...*options.DatabaseOptions) *mongo.Database
}

func NewClient(opts ...ClientOption) (*mongo.Client, error) {
	p := ClientParam{}
	for _, opt := range opts {
		opt(&p)
	}

	if !p.ModeSupported() {
		return nil, fmt.Errorf("mode is not supported")
	}
	if !p.AuthSupported() {
		return nil, fmt.Errorf("auth is not supported")
	}

	mongoOption := options.Client()

	if p.ModeStandalone() {
		mongoOption.SetHosts([]string{fmt.Sprintf("%s:%d", p.StdHost, p.StdPort)})
	} else if p.ModeReplication() {
		mongoOption.
			SetHosts(p.RsHosts).
			SetReplicaSet(p.RsName).
			SetReadPreference(readpref.Secondary())
	}

	if p.AuthBasic() {
		mongoOption.SetAuth(options.Credential{
			Username:   p.AuthUser,
			Password:   p.AuthPassword,
			AuthSource: p.AuthSource,
		})
	}

	return mongo.NewClient(mongoOption)
}
