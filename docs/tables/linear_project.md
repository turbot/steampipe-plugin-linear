---
title: "Steampipe Table: linear_project - Query Linear Projects using SQL"
description: "Allows users to query Linear Projects. This table returns information about the projects in Linear, including their name, status, and description."
---

# Table: linear_project - Query Linear Projects using SQL

Linear Project is a resource in Linear, a software development management tool. It allows users to organize tasks, issues, and objectives into specific projects. Each project in Linear has a unique identifier, name, and status, and may contain additional details such as description, label, and team information.

## Table Usage Guide

The `linear_project` table provides insights into projects within Linear. As a project manager or a software development team member, you can explore project-specific details through this table, including their status, description, and associated team information. Utilize it to manage and monitor the progress of different projects, understand their scope, and ensure timely completion of tasks and objectives.

## Examples

### Basic info
Discover the segments that are currently in progress within your projects and assess their status and timeline. This can help you understand the overall progress and manage your projects more efficiently.

```sql+postgres
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

```sql+sqlite
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
Explore which projects are in the planning stage to effectively manage resources and prioritize tasks. This aids in strategic decision-making by providing insights into upcoming projects.

```sql+postgres
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

```sql+sqlite
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
Explore which projects are still ongoing. This is useful for tracking progress and identifying tasks that may need additional resources or attention.

```sql+postgres
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

```sql+sqlite
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
Discover the projects that were established by administrative users. This can be particularly useful for auditing or understanding the distribution of project creation responsibilities within your team.

```sql+postgres
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

```sql+sqlite
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
  json_extract(creator, '$.admin') = 'true';
```

### Show lead details of each project
Explore the leadership details of each project to gain insights into their activity status and roles. This can help in assessing the active involvement and administrative roles of the leads in different projects.

```sql+postgres
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

```sql+sqlite
select
  json_extract(lead, '$.id') as lead_id,
  json_extract(lead, '$.name') as name,
  json_extract(lead, '$.email') as email,
  json_extract(lead, '$.active') as active,
  json_extract(lead, '$.admin') as admin,
  json_extract(lead, '$.statusLabel') as status_label
from
  linear_project;
```