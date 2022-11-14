package main

import (
	"google.golang.org/grpc"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:4467", grpc.WithInsecure())
	if err != nil {
		panic("Encountered error: " + err.Error())
	}

	writeClient := acl.NewWriteServiceClient(conn)

	// _, err = client.TransactRelationTuples(context.Background() ...
	conn, err = grpc.Dial("127.0.0.1:4466", grpc.WithInsecure())

	readClient := acl.NewReadServiceClient(conn)
	// _, err = readClient.ListRelationTuples(context.Background()...

	conn, err = grpc.Dial("127.0.0.1:4466", grpc.WithInsecure())
	checkClient := acl.NewCheckServiceClient(conn)
	// _, err = checkClient.Check(context.Background() ...

	conn, err = grpc.Dial("127.0.0.1:4466", grpc.WithInsecure())
	expandClient := acl.NewExpandServiceClient(conn)
	// _, err = expandClient.Expand(context.Background() ...
}
