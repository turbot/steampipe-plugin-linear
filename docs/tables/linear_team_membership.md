# Table: linear_team_membership

Linear team membership is a feature of the Linear project management tool that allows you to add users to your team and give them access to your workspace. Team members can then view and edit issues, tasks, and conversations, and they can also participate in team discussions.

## Examples

### Basic info

```sql
select
  id,
  created_at,
  owner,
  sort_order,
  updated_at
from
  linear_team_membership;
```

### List teams with owner details

```sql
select
  id,
  jsonb_pretty(team) as team,
  jsonb_pretty(membership_user) as user,
  updated_at
from
  linear_team_membership
where
  owner;
```

### List members of a particular team

```sql
select
  user ->> 'id' as user_id,
  user ->> 'name' as name,,
  user ->> 'admin' as admin,
  user ->> 'email' as email,
  user ->> 'active' as active
from
  linear_team_membership
where
  team ->> 'name' = 'linear_team';
```

### List archived membership

```sql
select
  id,
  created_at,
  owner,
  sort_order,
  updated_at
from
  linear_team_membership
where
  archived_at is not null;
```
