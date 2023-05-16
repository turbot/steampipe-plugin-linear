# Table: linear_comment

Linear comments are a simple and effective way to add documentation to your code. They can help other developers understand what your code is doing, and they can help you to keep track of changes that have been made to your code.

## Examples

### Basic info

```sql
select
  id,
  title,
  created_at,
  edited_at,
  updated_at,
  url
from
  linear_comment;
```

### Show user details of each comment

```sql
select
  id,
  title,
  creator ->> 'id' as creator_id,
  creator ->> 'name' as creator_name,
  creator ->> 'active' as active,
  creator ->> 'email' as email,
  creator ->> 'admin' as admin,
  creator ->> 'createdAt' as created_at
from
  linear_comment;
```

### List comments for a particular issue

```sql
select
  id,
  title,
  created_at,
  edited_at,
  updated_at,
  url
from
  linear_comment
where
  issue ->> 'title' = 'attachment check';
```

### List comments written by admin

```sql
select
  id,
  title,
  created_at,
  edited_at,
  updated_at,
  url
from
  linear_comment
where
  creator ->> 'admin' = 'true';
```

### List comments older than 90 days

```sql
select
  id,
  title,
  created_at,
  edited_at,
  updated_at,
  url
from
  linear_comment
where
  created_at > now() - interval '90' day;
```
