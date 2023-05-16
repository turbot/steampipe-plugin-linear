# Table: linear_organization

The organization or workspace is your home in Linear and where all issues and interactions live. Within your workspace, you'll create teams to group people who work together, create issues in specific teams, and use cycles and projects to define sets of work.

## Examples

### Basic info

```sql
select
  id,
  title,
  created_at,
  url_key,
  user_count,
  roadmap_enabled,
  release_channel,
  updated_at
from
  linear_organization;
```

### List teams in the organization

```sql
select
  t.id,
  t.title,
  t.color,
  t.key,
  t.private,
  t.updated_at
from
  linear_team as t,
  linear_organization as o
where
  o.id = t.organization ->> 'id';
```

### List users in the organization

```sql
select
  u.id,
  u.title,
  u.active,
  u.admin,
  u.email,
  u.updated_at
from
  linear_user as u,
  linear_organization as o
where
  o.id = u.organization ->> 'id';
```

### List integrations in the organization

```sql
select
  i.id,
  i.created_at,
  i.service,
  i.updated_at
from
  linear_integration as i,
  linear_organization as o
where
  o.id = i.organization ->> 'id';
```

### Show subscription details

```sql
select
  subscription ->> 'id' as creator_id,
  subscription ->> 'nextBillingAt' as next_billing_at,
  subscription ->> 'seats' as seats,
  subscription ->> 'type' as type,
  subscription ->> 'createdAt' as created_at
from
  linear_organization;
```