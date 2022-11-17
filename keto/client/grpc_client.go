//based on ory's keto cli's grpc client:
//https://github.com/ory/keto/blob/6c0e1ba87f4d3a355cebd0ea77f28319be2dd606/cmd/client/grpc_client.go

package client

import (
	"context"
	"crypto/tls"
	"fmt"
	"os"
	"strings"
	"time"

	"golang.org/x/oauth2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/credentials/oauth"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

type contextKeys string

const (
	ReadRemoteDefault  = "127.0.0.1:4466"
	WriteRemoteDefault = "127.0.0.1:4467"
	EnvReadRemote      = "KETO_READ_REMOTE"
	EnvWriteRemote     = "KETO_WRITE_REMOTE"
	EnvAuthToken       = "KETO_BEARER_TOKEN" // nosec G101 -- just the key, not the value
	EnvAuthority       = "KETO_AUTHORITY"

	ContextKeyTimeout contextKeys = "timeout"
)

type connectionDetails struct {
	token, authority     string
	skipHostVerification bool
	noTransportSecurity  bool
}

func (d *connectionDetails) dialOptions() (opts []grpc.DialOption) {
	if d.token != "" {
		opts = append(opts,
			grpc.WithPerRPCCredentials(
				oauth.NewOauthAccess(&oauth2.Token{AccessToken: d.token})))
	}
	if d.authority != "" {
		opts = append(opts, grpc.WithAuthority(d.authority))
	}

	// TLS settings
	switch {
	case d.noTransportSecurity:
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	case d.skipHostVerification:
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			// nolint explicity set through scary flag
			InsecureSkipVerify: true,
		})))
	default:
		// Defaults to the default host root CA bundle
		opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(nil)))
	}
	return opts
}

func getRemote(envRemote, remoteDefault string) (remote string) {
	defer (func() {
		if strings.HasPrefix(remote, "http://") || strings.HasPrefix(remote, "https://") {
			_, _ = fmt.Fprintf(os.Stderr, "remote \"%s\" seems to be an http URL instead of a remote address\n", remote)
		}
	})()

	if remote, isSet := os.LookupEnv(envRemote); isSet {
		return remote
	} else {
		_, _ = fmt.Fprintf(os.Stderr, "env var %s is not set, falling back to %s\n", envRemote, remoteDefault)
		return remoteDefault
	}
}

func getAuthority() string {
	return os.Getenv(EnvAuthority)
}

func getConnectionDetails() connectionDetails {
	return connectionDetails{
		token:                os.Getenv(EnvAuthToken),
		authority:            getAuthority(),
		skipHostVerification: false,
		noTransportSecurity:  false,
	}
}

func GetReadConn(ctx context.Context) (*grpc.ClientConn, error) {
	return Conn(ctx,
		getRemote(EnvReadRemote, ReadRemoteDefault),
		getConnectionDetails(),
	)
}

func GetWriteConn(ctx context.Context) (*grpc.ClientConn, error) {
	return Conn(ctx,
		getRemote(EnvWriteRemote, WriteRemoteDefault),
		getConnectionDetails(),
	)
}

func Conn(ctx context.Context, remote string, details connectionDetails) (*grpc.ClientConn, error) {
	timeout := 3 * time.Second
	if d, ok := ctx.Value(ContextKeyTimeout).(time.Duration); ok {
		timeout = d
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	return grpc.DialContext(
		ctx,
		remote,
		append([]grpc.DialOption{
			grpc.WithBlock(),
			grpc.WithDisableHealthCheck(),
		}, details.dialOptions()...)...,
	)
}

type WriteServiceClientForDummies interface {
	rts.WriteServiceClient
	Insert(ctx context.Context, tuples []*rts.RelationTuple) (*rts.TransactRelationTuplesResponse, error)
	Delete(ctx context.Context, tuples []*rts.RelationTuple) (*rts.DeleteRelationTuplesResponse, error)
}

type WriteClient struct {
	wsc rts.WriteServiceClient
}

func (client *WriteClient) Insert(ctx context.Context, tuples []*rts.RelationTuple) (*rts.TransactRelationTuplesResponse, error) {
	return client.wsc.TransactRelationTuples(ctx, &rts.TransactRelationTuplesRequest{
		RelationTupleDeltas: rts.RelationTupleToDeltas(tuples, rts.RelationTupleDelta_ACTION_INSERT),
	})
}

func (client *WriteClient) Delete(ctx context.Context, tuples []*rts.RelationTuple) (*rts.TransactRelationTuplesResponse, error) {
	return client.wsc.TransactRelationTuples(ctx, &rts.TransactRelationTuplesRequest{
		RelationTupleDeltas: rts.RelationTupleToDeltas(tuples, rts.RelationTupleDelta_ACTION_INSERT),
	})
}

func NewWriteClient(cc grpc.ClientConnInterface) WriteClient {
	return WriteClient{
		wsc: rts.NewWriteServiceClient(cc),
	}
}
