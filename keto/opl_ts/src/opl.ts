import { Namespace, SubjectSet, Context } from "@ory/keto-namespace-types"

class User implements Namespace { }

class Group implements Namespace {
  related: {
    member: (User | Group)[]
  }
}

class Project implements Namespace {
  related: {
    access: (User | Group)[]
  }
}

class Role implements Namespace {
  related: {
    principal: Project[]
  }
  permits = {
    can_assume: (ctx: Context) => this.related.principal.traverse((p) => p.related.access.includes(ctx.subject))
  }
}

class Policy implements Namespace {
  related: {
    allow: Role[]
  }
  permits = {
    allow: (ctx: Context) => this.related.allow.traverse((r) => r.permits.can_assume(ctx))
  }
}

class ResourcePolicy implements Namespace {
  related: {
    trust: (User | Group)[]
  }
  permits = {
    allow: (ctx: Context) => this.related.trust.includes(ctx.subject)
  }
}

class KubernetesResourceType implements Namespace {
  related: {
    create: Policy[]
    delete: Policy[]
    deletecollection: Policy[]
    get: Policy[]
    list: Policy[]
    patch: Policy[]
    update: Policy[]
    watch: Policy[]
  }

  permits = {
    can_create: (ctx: Context) => this.related.create.traverse((p) => p.permits.allow(ctx)),
    can_delete: (ctx: Context) => this.related.delete.traverse((p) => p.permits.allow(ctx)),
    can_deletecollection: (ctx: Context) => this.related.delete.traverse((p) => p.permits.allow(ctx)),
    can_get: (ctx: Context) => this.related.delete.traverse((p) => p.permits.allow(ctx)),
    can_list: (ctx: Context) => this.related.delete.traverse((p) => p.permits.allow(ctx)),
    can_patch: (ctx: Context) => this.related.delete.traverse((p) => p.permits.allow(ctx)),
    can_update: (ctx: Context) => this.related.delete.traverse((p) => p.permits.allow(ctx)),
    can_watch: (ctx: Context) => this.related.delete.traverse((p) => p.permits.allow(ctx)),
  }
}

class KubricksResourceType implements Namespace {
  related: {
    apiaccess: Policy[]
  }

  permits = {
    can_apiaccess: (ctx: Context) => this.related.apiaccess.traverse((p) => p.permits.allow(ctx)),
  }
}

class ServiceResource implements Namespace {
  related: {
    owner: User[]
    k8s_instance: KubernetesResourceType[]
    kbrx_instance: KubricksResourceType[]
    get: ResourcePolicy[]
    watch: ResourcePolicy[]
    apiaccess: (ResourcePolicy | Policy)[]
  }
  permits = {
    // Kubernetes rule verbs
    can_delete: (ctx: Context) => this.related.k8s_instance.traverse((i) => i.permits.can_delete(ctx)) || this.related.owner.includes(ctx.subject),
    can_get: (ctx: Context) => this.related.k8s_instance.traverse((i) => i.permits.can_get(ctx)) || this.related.owner.includes(ctx.subject),
    can_patch: (ctx: Context) => this.related.k8s_instance.traverse((i) => i.permits.can_patch(ctx)) || this.related.owner.includes(ctx.subject),
    can_update: (ctx: Context) => this.related.k8s_instance.traverse((i) => i.permits.can_update(ctx)) || this.related.owner.includes(ctx.subject),
    can_watch: (ctx: Context) => this.related.k8s_instance.traverse((i) => i.permits.can_watch(ctx)) || this.related.owner.includes(ctx.subject),
    // Kubricks rule verbs
    can_apiaccess: (ctx: Context) => this.related.kbrx_instance.traverse((i) => i.permits.can_apiaccess(ctx)) || this.related.apiaccess.traverse((p) => p.permits.allow(ctx)) || this.related.owner.includes(ctx.subject),
  }
}
