package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	acl "github.com/ory/keto/proto/ory/keto/acl/v1alpha1"
)

func main() {
	conn, err := grpc.Dial("bold-dubinsky-wgsl9ep42v.projects.oryapis.com:443")
	if err != nil {
		panic("Encountered error: " + err.Error())
	}

	client := acl.NewWriteServiceClient(conn)

	_, err = client.TransactRelationTuples(context.Background(), &acl.TransactRelationTuplesRequest{
		RelationTupleDeltas: []*acl.RelationTupleDelta{
			{
				Action: acl.RelationTupleDelta_INSERT,
				RelationTuple: &acl.RelationTuple{
					Namespace: "blog_posts",
					Object:    "my-first-blog-post",
					Relation:  "read",
					Subject: &acl.Subject{Ref: &acl.Subject_Id{
						Id: "alice",
					}},
				},
			},
		},
	})
	if err != nil {
		panic("Encountered error: " + err.Error())
	}

	fmt.Println("Successfully created tuple")
	readConn, err := grpc.Dial("bold-dubinsky-wgsl9ep42v.projects.oryapis.com:443")
	if err != nil {
		panic("Encountered error: " + err.Error())
	}
	checkClient := acl.NewCheckServiceClient(readConn)

	check, err := checkClient.Check(context.Background(), &acl.CheckRequest{
		Namespace: "blog_posts",
		Object:    "my-first-blog-post",
		Relation:  "read",
		Subject: &acl.Subject{Ref: &acl.Subject_Id{
			Id: "user1",
		}},
	})

	if check.Allowed {
		fmt.Println("Alice has access to my-first-blog-post")
	}

}
