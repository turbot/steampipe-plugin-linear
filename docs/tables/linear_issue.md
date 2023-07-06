# Table: linear_issue

The most basic concept in Linear is the issue. It's the building block of the app and most concepts in Linear are either associated with issues or group issues together. Issues are always tied to a specific team.

## Examples

### Basic info

```sql
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue;
```

### List urgent issues

```sql
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  priority = 1;
```

### List issues that have not been started

```sql
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  started_at is null;
```

### List issues that have crossed the due date

```sql
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  due_date < now();
```

### List trashed issues

```sql
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  trashed;
```

### List issues created by admin

```sql
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  creator ->> 'admin' = 'true';
```

### List issues of a particular team

```sql
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  team ->> 'name' = 'linear';
```

### List unassigned issues

```sql
select
  id,
  title,
  created_at,
  branch_name,
  priority,
  estimate,
  updated_at
from
  linear_issue
where
  assignee is null;
```
