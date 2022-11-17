package opl_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	//. "github.com/onsi/gomega"
	"github.com/ory/x/pointerx"

	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

var _ = Describe("Verify expected behaviour of the opl configuration.", func() {
	var _ = Describe("Scenario to cover most constellations.", func() {
		var _ = BeforeEach(func() {
			wcl.DeleteRelationTuples(context.TODO(), &rts.DeleteRelationTuplesRequest{})
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
			It("should be able to insert", func() {
				getStringPtr := func(s string) *string {
					return pointerx.Ptr(s)
				}
				query := rts.RelationQuery{
					Namespace: getStringPtr("User"),
					Object:    getStringPtr(""),
					Relation:  getStringPtr(""),
				}
				_, err := wcl.DeleteRelationTuples(context.TODO(), &rts.DeleteRelationTuplesRequest{
					RelationQuery: &query,
				})
				if err != nil {
					panic("Encountered error: " + err.Error())
				}
				_ = tuples
				//_, err = wcl.TransactRelationTuples(context.Background(), &rts.TransactRelationTuplesRequest{
				//	RelationTupleDeltas: rts.RelationTupleToDeltas(tuples, rts.RelationTupleDelta_ACTION_INSERT),
				//})
				//if err != nil {
				//	panic("Encountered error: " + err.Error())
				//}

			})
		})
	})
})
