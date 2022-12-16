package opl_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	//. "github.com/onsi/gomega"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
	px "github.com/ory/x/pointerx"

	"github.com/rauerhans/laborat-ory/keto/client"
)

var _ = Describe("Verify expected behaviour of the opl configuration.", func() {
	var _ = Describe("Scenario to cover most constellations.", func() {
		//5 users: Hans, David, Nico, Lianet, Sophie
		//2 groups: AllUsers, Ops
		//1 project: Manhattan
		//2 roles in project Manhattan: creator, editor
		//2 policies:
		BeforeEach(func() {
			//set up database before each test
			err := kcl.CreateTuples(context.Background(), scenario_1)
			if err != nil {
				panic("Encountered error: " + err.Error())
			}

		})
		AfterEach(func() {
			//tear down database entries after each test
			query := rts.RelationQuery{
				Namespace: nil,
				Object:    nil,
				Relation:  nil,
				Subject:   nil,
			}
			err := kcl.DeleteAllTuples(context.Background(), &query)
			if err != nil {
				panic("Encountered error: " + err.Error())
			}
		})
		It("should be able to list all users", func() {
			query := rts.RelationQuery{
				Namespace: px.Ptr("Group"),
				Object:    px.Ptr("AllUsers"),
				Relation:  nil,
				Subject:   nil,
			}

			respTuples, err := kcl.QueryAllTuples(context.Background(), &query, 100)
			if err != nil {
				panic("Encountered error: " + err.Error())
			}
			client.PrintTableFromRelationTuples(respTuples, GinkgoWriter)
		})
		It("Hans can create S3Resource", func() {
			//query := rts.RelationQuery{
			//	Namespace: px.Ptr("Group"),
			//	Object:    px.Ptr("AllUsers"),
			//	Relation:  nil,
			//	Subject:   nil,
			//}
			//resp, err := ccl.Check(context.Background(), &rts.CheckRequest{
			//	Tuple: &rts.RelationTuple{
			//		Namespace: "S3ResourceType",
			//		Object:    "S3",
			//		Relation:  ,
			//		Subject:   sub,
			//	},
			//	MaxDepth: 7,
			//})
			//resp, err := rcl.ListRelationTuples(context.Background(), &rts.ListRelationTuplesRequest{
			//	RelationQuery: &query,
			//})
			//if err != nil {
			//	panic("Encountered error: " + err.Error())
			//}
			//client.PrintTableFromRelationTuples(resp.RelationTuples, GinkgoWriter)
		})
	})
})
