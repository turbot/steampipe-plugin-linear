# Table: linear_project

Projects define larger pieces of work that have a clear outcome or completion date, such as launching a new feature. They can be shared across multiple teams and come with their own unique features, graph, and notification options.

## Examples

### Basic info

```sql
select
  id,
  title,
  created_at,
  color,
  progress,
  scope,
  state,
  updated_at
from
  linear_project;
```

### List planned projects

```sql
select
  id,
  title,
  created_at,
  color,
  progress,
  scope,
  state,
  updated_at
from
  linear_project
where
  state = 'planned';
```

### List projects which are incomplete

```sql
select
  id,
  title,
  created_at,
  color,
  progress,
  scope,
  state,
  updated_at
from
  linear_project
where
  completed_at is null;
```

### List projects created by admin

```sql
select
  id,
  title,
  created_at,
  color,
  progress,
  scope,
  state,
  updated_at
from
  linear_project
where
  creator ->> 'admin' = 'true';
```

### Show lead details of each project

```sql
select
  lead ->> 'id' as lead_id,
  lead ->> 'name' as name,
  lead ->> 'email' as email,
  lead ->> 'active' as active,
  lead ->> 'admin' as admin,
  lead ->> 'statusLabel' as status_label
from
  linear_project;
```
