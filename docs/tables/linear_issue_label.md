# Table: linear_issue_label

A label is a keyword or short phrase that you can use to categorize and organize your tasks, issues, and conversations. Labels can be used to create custom views of your work, and they can also be used to filter and search for specific items.

## Examples

### Basic info

```sql
select
  id,
  title,
  created_at,
  color,
  updated_at
from
  linear_issue_label;
```

### List labels which are not associated with any team

```sql
select
  id,
  title,
  created_at,
  color,
  updated_at
from
  linear_issue_label
where
  team is null;
```

### List all labels for each issues

```sql
select
  i.title as issue_title,
  l.title as label_tile
from
  linear_issue as i,
  linear_issue_label as l,
  jsonb_array_elements(issue_ids) as ids
where
  i.id = ids ->> 'id';
```

### Show archived labels

```sql
select
  id,
  title,
  created_at,
  color,
  updated_at
from
  linear_issue_label
where
  archived_at is not null;
```

### List labels created by admin

```sql
select
  id,
  title,
  created_at,
  color,
  updated_at
from
  linear_issue_label
where
  creator ->> 'admin' = 'true';
```

### List child labels of a particular label

```sql
select
  id,
  title,
  created_at,
  color,
  updated_at
from
  linear_issue_label
where
  parent ->> 'name' = 'issueLabel';
```
