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

	//-------- Create policies ---------
	// policies bundle permissions that are granted to a role
	
	// Policy: AdminPolicy
	{
		Namespace: "Policy",
		Object:    "AdminPolicy",
		Relation:  "allow",
		Subject: rts.NewSubjectSet(
			"Role",
			"Admin",
			"",
		),
	},
	// Policy: DevPolicy
	{
		Namespace: "Policy",
		Object:    "DevPolicy",
		Relation:  "allow",
		Subject: rts.NewSubjectSet(
			"Role",
			"Dev",
			"",
		)
	},

	//-------- create resource types and implicitly create permissions ---------
	// The resource types namespaces comprise all resources that are available in the system
	// They act as classes of resources for which you can create permissions by attaching them to policies
	// Any principal with a bound policy that has a permission for a certain resource type can perform the granted actions on all instances of resources of that type
	// It's also the only place where it makes sense to define permission relations for the creation of resources

	// KubernetesResourceType: Service
	// All Kubernetes primitive types (and maybe by extension CRDs, not sure yet) should be defined in the KubernetesResourceType namespace
	// For each action (rule verb) that is defined in the Kubernetes API for a resource type we create an according permission relation tuple

	// create permissions for the `Service` Kubernetes resource type and bundle them in the `AdminPolicy`
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

	// create certain narrow permissions for the `Service` Kubernetes resource type and bundle them in the `DevPolicy`
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
	
	// KubricksResourceType: MLFlow 
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
