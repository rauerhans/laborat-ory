package opl_test

import (
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

var scenario_1 = []*rts.RelationTuple{
	//-------- create Groups and Users ---------
	// Group: AllUsers
	// User objects are created implicitly through member relation tuples to a group `AllUsers` that contains all users
	// In a live Ory Kratos/Keto setup this group should reflect the users that are registered with Kratos
	{
		Namespace: "Group",
		Object:    "AllUsers",
		Relation:  "member",
		Subject: rts.NewSubjectSet(
			"User",
			"Hans",
			"",
		),
	},
	{
		Namespace: "Group",
		Object:    "AllUsers",
		Relation:  "member",
		Subject: rts.NewSubjectSet(
			"User",
			"David",
			"",
		),
	},
	{
		Namespace: "Group",
		Object:    "AllUsers",
		Relation:  "member",
		Subject: rts.NewSubjectSet(
			"User",
			"Nico",
			"",
		),
	},
	{
		Namespace: "Group",
		Object:    "AllUsers",
		Relation:  "member",
		Subject: rts.NewSubjectSet(
			"User",
			"Lianet",
			"",
		),
	},
	{
		Namespace: "Group",
		Object:    "AllUsers",
		Relation:  "member",
		Subject: rts.NewSubjectSet(
			"User",
			"Sophie",
			"",
		),
	},
	// Group: Ops
	// Some users are members of the group `Ops` that should contain users with admin access intended for administrative tasks
	{
		Namespace: "Group",
		Object:    "Ops",
		Relation:  "member",
		Subject: rts.NewSubjectSet(
			"User",
			"Hans",
			"",
		),
	},
	{
		Namespace: "Group",
		Object:    "Ops",
		Relation:  "member",
		Subject: rts.NewSubjectSet(
			"User",
			"David",
			"",
		),
	},
	//-------- create a project ---------
	// access does not entail any permissions, it is just a relation that's there to confirm that a user is registered with the project
	// if a user is not registered with a project they cannot perform any actions on resources inside the project even if there are still roles and policies in place that allow certain actions
	// that way we can for example revoke all access by severing the access to a project without deleting all policies or roles
	// Project: Manhattan

	// all members of group `Ops` have access to the project `Manhattan`
	{
		Namespace: "Project",
		Object:    "Manhattan",
		Relation:  "access",
		Subject: rts.NewSubjectSet(
			"Group",
			"Ops",
			"",
		),
	},
	// additionally Nico and Lianet have explicit access to the project `Manhattan`
	{
		Namespace: "Project",
		Object:    "Manhattan",
		Relation:  "access",
		Subject: rts.NewSubjectSet(
			"User",
			"Nico",
			"",
		),
	},
	{
		Namespace: "Project",
		Object:    "Manhattan",
		Relation:  "access",
		Subject: rts.NewSubjectSet(
			"User",
			"Lianet",
			"",
		),
	},

	//-------- create roles ---------
	// every action from a project member is performed on behalf of a principal
	// at this time only roles can act as principals

	// Role: Admin
	// we create a role `Admin` as a principal of project `Manhattan` that will get various wide ranging permissions through appropriate policies
	{
		Namespace: "Role",
		Object:    "Admin",
		Relation:  "principal",
		Subject: rts.NewSubjectSet(
			"Project",
			"Manhattan",
			"",
		),
	},
	// Role: Dev
	// we create a role `Dev` as a principal of project `Manhattan` that will get more narrow permissions
	{
		Namespace: "Role",
		Object:    "Dev",
		Relation:  "principal",
		Subject: rts.NewSubjectSet(
			"Project",
			"Manhattan",
			"",
		),
	},

	//-------- create  ---------
	{
		Namespace: "KubernetesResourceType",
		Object:    "Service",
		Relation:  "create",
		Subject: rts.NewSubjectSet(
			"Policy",
			"AdminPolicy",
			"",
		),
	},
	{
		Namespace: "KubernetesResourceType",
		Object:    "Service",
		Relation:  "delete",
		Subject: rts.NewSubjectSet(
			"Policy",
			"AdminPolicy",
			"",
		),
	},
	{
		Namespace: "KubernetesResourceType",
		Object:    "Service",
		Relation:  "get",
		Subject: rts.NewSubjectSet(
			"Policy",
			"AdminPolicy",
			"",
		),
	},
	{
		Namespace: "KubernetesResourceType",
		Object:    "Service",
		Relation:  "list",
		Subject: rts.NewSubjectSet(
			"Policy",
			"AdminPolicy",
			"",
		),
	},
	{
		Namespace: "KubernetesResourceType",
		Object:    "Service",
		Relation:  "update",
		Subject: rts.NewSubjectSet(
			"Policy",
			"AdminPolicy",
			"",
		),
	},
	{
		Namespace: "KubernetesResourceType",
		Object:    "Service",
		Relation:  "get",
		Subject: rts.NewSubjectSet(
			"Policy",
			"DevPolicy",
			"",
		),
	},
	{
		Namespace: "KubernetesResourceType",
		Object:    "Service",
		Relation:  "list",
		Subject: rts.NewSubjectSet(
			"Policy",
			"DevPolicy",
			"",
		),
	},
	{
		Namespace: "KubricksResourceType",
		Object:    "MLFlow",
		Relation:  "accessapi",
		Subject: rts.NewSubjectSet(
			"Policy",
			"AdminPolicy",
			"",
		),
	},
	{
		Namespace: "ServiceResource",
		Object:    "MLFlowInstance",
		Relation:  "owner",
		Subject: rts.NewSubjectSet(
			"User",
			"Hans",
			"",
		),
	},
	{
		Namespace: "ServiceResource",
		Object:    "MLFlowInstance",
		Relation:  "owner",
		Subject: rts.NewSubjectSet(
			"User",
			"Hans",
			"",
		),
	},
}
