//based on ory's keto cli's grpc client:
//https://github.com/ory/keto/blob/6c0e1ba87f4d3a355cebd0ea77f28319be2dd606/cmd/client/grpc_client.go

package client

import (
	"context"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
	"google.golang.org/grpc"
	grpcHealthV1 "google.golang.org/grpc/health/grpc_health_v1"
)

type Client interface {
	//TODO
	//queryNamespaces()
	transactTuples(ins []*rts.RelationTuple, del []*rts.RelationTuple)
	createTuple(r *rts.RelationTuple) error
	deleteTuple(r *rts.RelationTuple) error
	deleteAllTuples(q *rts.RelationQuery) error
	queryTuple(q *rts.RelationQuery, opts ...PaginationOptionSetter) (*rts.ListRelationTuplesResponse, error)
	queryAllTuples(q *rts.RelationQuery, pagesize int) ([]*rts.RelationTuple, error)
	check(r *rts.RelationTuple) (error, bool)
	expand(r *rts.SubjectSet, depth int) (error, *rts.SubjectTree)
	waitUntilLive()
}

type grpcClient struct {
	connDetails ConnectionDetails
	wc, rc, oc  *grpc.ClientConn
	ctx         context.Context
}

//func (g *grpcClient) queryNamespaces() {
//	return
//}

func (g *grpcClient) transactTuples(ins []*rts.RelationTuple, del []*rts.RelationTuple) error {
	c := rts.NewWriteServiceClient(g.wc)

	deltas := append(
		rts.RelationTupleToDeltas(ins, rts.RelationTupleDelta_ACTION_INSERT),
		rts.RelationTupleToDeltas(del, rts.RelationTupleDelta_ACTION_DELETE)...,
	)

	_, err := c.TransactRelationTuples(g.ctx, &rts.TransactRelationTuplesRequest{
		RelationTupleDeltas: deltas,
	})
	return err
}

func (g *grpcClient) createTuple(r *rts.RelationTuple) error {
	return g.transactTuples([]*rts.RelationTuple{r}, nil)
}

func (g *grpcClient) deleteTuple(r *rts.RelationTuple) error {
	return g.transactTuples(nil, []*rts.RelationTuple{r})
}

func (g *grpcClient) deleteAllTuples(q *rts.RelationQuery) error {
	c := rts.NewWriteServiceClient(g.wc)
	_, err := c.DeleteRelationTuples(g.ctx, &rts.DeleteRelationTuplesRequest{
		RelationQuery: q,
	})
	return err
}

type (
	PaginationOptions struct {
		Token string `json:"page_token"`
		Size  int    `json:"page_size"`
	}
	PaginationOptionSetter func(*PaginationOptions) *PaginationOptions
)

func WithToken(t string) PaginationOptionSetter {
	return func(opts *PaginationOptions) *PaginationOptions {
		opts.Token = t
		return opts
	}
}

func WithSize(size int) PaginationOptionSetter {
	return func(opts *PaginationOptions) *PaginationOptions {
		opts.Size = size
		return opts
	}
}

func GetPaginationOptions(modifiers ...PaginationOptionSetter) *PaginationOptions {
	opts := &PaginationOptions{}
	for _, f := range modifiers {
		opts = f(opts)
	}
	return opts
}

func (g *grpcClient) queryTuple(q *rts.RelationQuery, opts ...PaginationOptionSetter) (*rts.ListRelationTuplesResponse, error) {
	c := rts.NewReadServiceClient(g.rc)
	pagination := GetPaginationOptions(opts...)
	resp, err := c.ListRelationTuples(g.ctx, &rts.ListRelationTuplesRequest{
		RelationQuery: q,
		PageToken:     pagination.Token,
		PageSize:      int32(pagination.Size),
	})
	return resp, err
}

func (g *grpcClient) queryAllTuples(q *rts.RelationQuery, pagesize int) ([]*rts.RelationTuple, error) {
	tuples := make([]*rts.RelationTuple, 0)
	resp, err := g.queryTuple(q, WithSize(pagesize))
	for resp.NextPageToken != "" && err == nil {
		resp, err = g.queryTuple(q, WithToken(resp.NextPageToken), WithSize(pagesize))
		tuples = append(tuples, resp.RelationTuples...)
	}
	return tuples, err
}

func (g *grpcClient) check(r *rts.RelationTuple) (bool, error) {
	c := rts.NewCheckServiceClient(g.rc)

	req := &rts.CheckRequest{
		Tuple: r,
	}
	resp, err := c.Check(g.ctx, req)

	return resp.Allowed, err
}

func (g *grpcClient) expand(ss *rts.Subject, depth int) (*rts.SubjectTree, error) {
	c := rts.NewExpandServiceClient(g.rc)

	resp, err := c.Expand(g.ctx, &rts.ExpandRequest{
		Subject:  ss,
		MaxDepth: int32(depth),
	})
	return resp.Tree, err
}

//TODO not sure if this is the correct thing to do
func (g *grpcClient) waitUntilLive() error {
	c := grpcHealthV1.NewHealthClient(g.rc)

	ctx, cancel := context.WithCancel(g.ctx)
	defer cancel()

	cl, err := c.Watch(ctx, &grpcHealthV1.HealthCheckRequest{})
	if err != nil {
		return err
	}

	for {
		select {
		case <-g.ctx.Done():
			return nil
		default:
		}
		resp, err := cl.Recv()

		if resp.Status == grpcHealthV1.HealthCheckResponse_SERVING {
			return nil
		}
		if err != nil {
			return err
		}
	}
}
