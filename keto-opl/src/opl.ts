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
}

class S3ResourceType implements Namespace {
  related: {
    create: Policy[]
    write: Policy[]
    read: Policy[]
  }

  permits = {
    can_create: (ctx: Context) => this.related.create.traverse((p) => p.permits.allow(ctx)),
    can_write: (ctx: Context) => this.related.write.traverse((p) => p.permits.allow(ctx)),
    can_read: (ctx: Context) =>
      this.related.read.traverse((p) => p.permits.allow(ctx)) ||
      this.related.write.traverse((p) => p.permits.allow(ctx))
  }
}

class S3Resource implements Namespace {
  related: {
    instance: S3ResourceType[]
    write: ResourcePolicy[]
    read: ResourcePolicy[]
  }
  permits = {
    can_write: (ctx: Context) =>
      this.related.instance.traverse((i) => i.permits.can_write(ctx)) ||
      this.related.write.traverse((rp) => rp.related.trust.includes(ctx.subject)),
    can_read: (ctx: Context) =>
      this.related.instance.traverse((i) => i.permits.can_read(ctx)) || 
      this.related.read.traverse((rp) => rp.related.trust.includes(ctx.subject)) ||
      this.related.write.traverse((rp) => rp.related.trust.includes(ctx.subject))
  }
}
