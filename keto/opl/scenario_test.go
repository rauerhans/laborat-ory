package opl_test

import (
	rts "github.com/ory/keto/proto/ory/keto/relation_tuples/v1alpha2"
)

var scenario_1 = []*rts.RelationTuple{
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
	{
		Namespace: "S3ResourceType",
		Object:    "S3",
		Relation:  "create",
		Subject: rts.NewSubjectSet(
			"Policy",
			"CreatePolicy",
			"",
		),
	},
	{
		Namespace: "S3ResourceType",
		Object:    "S3",
		Relation:  "write",
		Subject: rts.NewSubjectSet(
			"Policy",
			"EditPolicy",
			"",
		),
	},
	{
		Namespace: "S3ResourceType",
		Object:    "S3",
		Relation:  "read",
		Subject: rts.NewSubjectSet(
			"Policy",
			"EditPolicy",
			"",
		),
	},
	{
		Namespace: "Policy",
		Object:    "CreatePolicy",
		Relation:  "allow",
		Subject: rts.NewSubjectSet(
			"Role",
			"Admin",
			"",
		),
	},
	{
		Namespace: "Policy",
		Object:    "EditPolicy",
		Relation:  "allow",
		Subject: rts.NewSubjectSet(
			"Role",
			"Admin",
			"",
		),
	},
	{
		Namespace: "Policy",
		Object:    "EditPolicy",
		Relation:  "allow",
		Subject: rts.NewSubjectSet(
			"Role",
			"Dev",
			"",
		),
	},
	{
		Namespace: "Policy",
		Object:    "EditPolicy",
		Relation:  "allow",
		Subject: rts.NewSubjectSet(
			"Role",
			"Dev",
			"",
		),
	},
	{
		Namespace: "S3Resource",
		Object:    "XFiles",
		Relation:  "instance",
		Subject: rts.NewSubjectSet(
			"S3ResourceType",
			"S3",
			"",
		),
	},
	{
		Namespace: "S3Resource",
		Object:    "XFiles",
		Relation:  "read",
		Subject: rts.NewSubjectSet(
			"ResourcePolicy",
			"XFilesReadOnly",
			"",
		),
	},
	{
		Namespace: "ResourcePolicy",
		Object:    "XFilesReadOnly",
		Relation:  "trust",
		Subject: rts.NewSubjectSet(
			"User",
			"Sophie",
			"",
		),
	},
}
