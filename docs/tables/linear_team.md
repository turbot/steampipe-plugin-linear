# Table: linear_team

A team in the Linear app is a group of users who are working together on a project. Teams can be created by admins and can be public or private. Public teams are visible to all users in the workspace, while private teams are only visible to team members.

## Examples

### Basic info

```sql
select
  id,
  title,
  created_at,
  color,
  private,
  triage_enabled,
  key,
  default_issue_estimate,
  auto_close_period,
  updated_at
from
  linear_team;
```

### List public teams

```sql
select
  id,
  title,
  created_at,
  color,
  private,
  triage_enabled,
  key,
  default_issue_estimate,
  auto_close_period,
  updated_at
from
  linear_team
where
  not private;
```

### List teams that are not using cycles

```sql
select
  id,
  title,
  created_at,
  color,
  private,
  triage_enabled,
  key,
  default_issue_estimate,
  auto_close_period,
  updated_at
from
  linear_team
where
  not cycles_enabled;
```

### List teams that are not associated with any integration

```sql
select
  id,
  title,
  created_at,
  color,
  private,
  triage_enabled,
  key,
  default_issue_estimate,
  auto_close_period,
  updated_at
from
  linear_team
where
  integrations_settings is null;
```

### List archived teams

```sql
select
  id,
  title,
  created_at,
  color,
  private,
  triage_enabled,
  key,
  default_issue_estimate,
  auto_close_period,
  updated_at
from
  linear_team
where
  archived_at is not null;
```

### List teams where triage mode is enabled

```sql
select
  id,
  title,
  created_at,
  color,
  private,
  key,
  default_issue_estimate,
  auto_close_period,
  updated_at
from
  linear_team
where
  triage_enabled;
```
