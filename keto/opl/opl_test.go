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
				Relation:  px.Ptr("usermember"),
				Subject:   nil,
			}

			respTuples, err := kcl.QueryAllTuples(context.Background(), &query, 100)
			if err != nil {
				panic("Encountered error: " + err.Error())
			}
			client.PrintTableFromRelationTuples(respTuples, GinkgoWriter)
		})
		It("Group `Ops` and by extension `Hans` and `David` can act as principals of project Manhattan", func() {
			query := rts.RelationTuple{
				Namespace: "Role",
				Object:    "Admin",
				Relation:  "can_assume",
				Subject: rts.NewSubjectSet(
					"Group",
					"Ops",
					"",
				),
			}
			ok, err := kcl.Check(context.Background(), &query)
			if err != nil {
				panic("Encountered error: " + err.Error())
			}
			GinkgoWriter.Printf("Group `Ops` can assume role `Admin`: %v\n", ok)

			for _, user := range []string{"Hans", "David"} {
				query = rts.RelationTuple{
					Namespace: "Role",
					Object:    "Admin",
					Relation:  "can_assume",
					Subject: rts.NewSubjectSet(
						"User",
						user,
						"",
					),
				}
				ok, err = kcl.Check(context.Background(), &query)
				if err != nil {
					panic("Encountered error: " + err.Error())
				}
				GinkgoWriter.Printf("User `%v` can assume role `Admin`: %v\n", user, ok)

			}
		})
		It("Groups `Ops` members can create, delete, get, list Kubernetes Secrets, because they can assume the Admin role", func() {
			for _, user := range []string{"Hans", "David"} {
				for _, action := range []string{"create", "delete", "get", "list"} {
					query := rts.RelationTuple{
						Namespace: "KubernetesResourceType",
						Object:    "Secret",
						// help me build a string format
						Relation: "can_" + action,
						Subject: rts.NewSubjectSet(
							"User",
							user,
							"",
						),
					}
					ok, err := kcl.Check(context.Background(), &query)
					if err != nil {
						panic("Encountered error: " + err.Error())
					}
					GinkgoWriter.Printf("User `%v` can %v Secret: %v\n", user, action, ok)
				}
			}
		})
	})
})
