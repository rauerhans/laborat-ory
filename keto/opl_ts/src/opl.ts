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
    create: Policy[]
    delete: Policy[]
    get: Policy[]
    list: Policy[]
    update: Policy[]

    accessapi: Policy[]

    hassecret: KubernetesResourceType[]
    setsecret: Policy[]
    getsecret: Policy[]
  }

  permits = {
    // do things with the resource definition and its properties
    can_create: (ctx: Context) => this.related.create.traverse((p) => p.permits.allow(ctx)),
    can_delete: (ctx: Context) => this.related.delete.traverse((p) => p.permits.allow(ctx)),
    can_get: (ctx: Context) => this.related.delete.traverse((p) => p.permits.allow(ctx)),
    can_list: (ctx: Context) => this.related.delete.traverse((p) => p.permits.allow(ctx)),
    can_update: (ctx: Context) => this.related.delete.traverse((p) => p.permits.allow(ctx)),

    // do things with the resources possible API
    can_accessapi: (ctx: Context) => this.related.accessapi.traverse((p) => p.permits.allow(ctx)),

    // do things with the resources possible kubernetes dependencies without giving broader access to the resource itself
    can_setsecret: (ctx: Context) => this.related.setsecret.traverse((p) => p.permits.allow(ctx)) || this.related.hassecret.traverse((p) => p.permits.can_delete(ctx) && p.permits.can_create(ctx)),
    can_getsecret: (ctx: Context) => this.related.getsecret.traverse((p) => p.permits.allow(ctx)) || this.related.hassecret.traverse((p) => p.permits.can_get(ctx)),
  }
}

class KubricksResource implements Namespace {
  related: {
    owner: User[]
    kbrx_instance: KubricksResourceType[]
    accessapi: (Policy | ResourcePolicy)[]
    setsecret: (Policy | ResourcePolicy)[]
    getsecret: (Policy | ResourcePolicy)[]
  }
  permits = {
    can_delete: (ctx: Context) => this.related.kbrx_instance.traverse((i) => i.permits.can_delete(ctx)) || this.related.owner.includes(ctx.subject),
    can_get: (ctx: Context) => this.related.kbrx_instance.traverse((i) => i.permits.can_get(ctx)) || this.related.owner.includes(ctx.subject),
    can_update: (ctx: Context) => this.related.kbrx_instance.traverse((i) => i.permits.can_update(ctx)) || this.related.owner.includes(ctx.subject),
    can_accessapi: (ctx: Context) => this.related.kbrx_instance.traverse((i) => i.permits.can_accessapi(ctx)) || this.related.accessapi.traverse((p) => p.permits.allow(ctx)) || this.related.owner.includes(ctx.subject),
    can_setsecret: (ctx: Context) => this.related.kbrx_instance.traverse((i) => i.permits.can_setsecret(ctx)) || this.related.setsecret.traverse((p) => p.permits.allow(ctx)) || this.related.owner.includes(ctx.subject),
    can_getsecret: (ctx: Context) => this.related.kbrx_instance.traverse((i) => i.permits.can_getsecret(ctx)) || this.related.getsecret.traverse((p) => p.permits.allow(ctx)) || this.related.owner.includes(ctx.subject),
  }
}
