package opl_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	//. "github.com/onsi/gomega"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"

	"github.com/rauerhans/laborat-ory/keto/client"
)

var _ = Describe("Verify expected behaviour of the opl configuration.", func() {
	var _ = Describe("Scenario to cover most constellations.", func() {
		BeforeEach(func() {
			tuples := []*rts.RelationTuple{
				{
					Namespace: "Group",
					Object:    "ops",
					Relation:  "member",
					Subject: rts.NewSubjectSet(
						"User",
						"Hans",
						"",
					),
				}, {
					Namespace: "Group",
					Object:    "ops",
					Relation:  "member",
					Subject: rts.NewSubjectSet(
						"User",
						"David",
						"",
					),
				}, {
					Namespace: "Group",
					Object:    "ops",
					Relation:  "member",
					Subject: rts.NewSubjectSet(
						"User",
						"Sophie",
						"",
					),
				},
			}
			_, err := wcl.TransactRelationTuples(context.TODO(), &rts.TransactRelationTuplesRequest{
				RelationTupleDeltas: rts.RelationTupleToDeltas(tuples, rts.RelationTupleDelta_ACTION_INSERT),
			})
			if err != nil {
				panic("Encountered error: " + err.Error())
			}

		})
		AfterEach(func() {
			query := rts.RelationQuery{
				Namespace: nil,
				Object:    nil,
				Relation:  nil,
				Subject:   nil,
			}
			_, err := wcl.DeleteRelationTuples(context.TODO(), &rts.DeleteRelationTuplesRequest{
				RelationQuery: &query,
			})
			if err != nil {
				panic("Encountered error: " + err.Error())
			}
		})

		It("should be able to delete all", func() {

		})
		It("should be able to list", func() {
			query := rts.RelationQuery{
				Namespace: nil,
				Object:    nil,
				Relation:  nil,
				Subject:   nil,

				//Namespace: pointerx.Ptr("User"),
			}
			resp, err := rcl.ListRelationTuples(context.Background(), &rts.ListRelationTuplesRequest{
				RelationQuery: &query,
			})
			if err != nil {
				panic("Encountered error: " + err.Error())
			}

			relationTuples, err := client.NewCollection(resp.RelationTuples)
			client.PrintTable(client.FormatTable, GinkgoWriter, relationTuples)
			if err != nil {
				panic("Encountered error: " + err.Error())
			}

		})
	})
})
