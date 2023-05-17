# Table: linear_user

In Linear, a user is an individual who has access to a Linear workspace. Users can create and edit issues, tasks, and conversations, and they can also participate in team discussions.

There are three types of users in Linear:

Admins: Admins have full access to a workspace and can manage team members.
Members: Members have read-write access to a workspace and can participate in team discussions.
Guests: Guests have read-only access to a workspace and cannot participate in team discussions.

## Examples

### Basic info

```sql
select
  id,
  title,
  active,
  admin,
  created_at,
  email,
  status_emoji,
  status_label,
  updated_at
from
  linear_user;
```

### List users who are admins

```sql
select
  id,
  title,
  active,
  admin,
  created_at,
  email,
  status_emoji,
  status_label,
  updated_at
from
  linear_user
where
  admin;
```

### List inactive users

```sql
select
  id,
  title,
  active,
  admin,
  created_at,
  email,
  status_emoji,
  status_label,
  updated_at
from
  linear_user
where
  not active;
```

### Show details of the currently authenticated user

```sql
select
  id,
  title,
  active,
  admin,
  created_at,
  email,
  status_emoji,
  status_label,
  updated_at
from
  linear_user
where
  is_me;
```

### List guest users

```sql
select
  id,
  title,
  active,
  admin,
  created_at,
  email,
  status_emoji,
  status_label,
  updated_at
from
  linear_user
where
  guest;
```

### List archived users

```sql
select
  id,
  title,
  active,
  admin,
  created_at,
  email,
  status_emoji,
  status_label,
  updated_at
from
  linear_user
where
  archived_at is not null;
```
